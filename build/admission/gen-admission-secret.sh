#!/bin/bash

set -e

SERVICE=${SERVICE:-"kubeedge-admission-service"}
SECRET=${SECRET:-"kubeedge-admission-secret"}
NAMESPACE=${NAMESPACE:-kubeedge}
CERTDIR=${CERTDIR:-"/etc/kubeedge/admission"}
ENABLE_CREATE_SECRET=${ENABLE_CREATE_SECRET:-true}

if [[ ! -x "$(command -v openssl)" ]]; then
    echo "openssl not found"
    exit 1
fi

csrName=${SERVICE}.${NAMESPACE}
mkdir -p ${SERVICE}
echo "creating certs in dir ${CERTDIR} "

cat <<EOF > ${CERTDIR}/csr.conf
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = ${SERVICE}
DNS.2 = ${SERVICE}.${NAMESPACE}
DNS.3 = ${SERVICE}.${NAMESPACE}.svc
EOF

openssl genrsa -out ${CERTDIR}/server-key.pem 2048
openssl req -new -key ${CERTDIR}/server-key.pem -subj "/CN=${SERVICE}.${NAMESPACE}.svc" -out ${CERTDIR}/server.csr -config ${CERTDIR}/csr.conf

# clean-up any previously created CSR for our service. Ignore errors if not present.
kubectl delete csr ${csrName} 2>/dev/null || true

# create  server cert/key CSR and  send to k8s API
cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${csrName}
spec:
  groups:
  - system:authenticated
  request: $(cat ${CERTDIR}/server.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - server auth
EOF

# verify CSR has been created
while true; do
    kubectl get csr ${csrName}
    if [ "$?" -eq 0 ]; then
        break
    fi
done

# approve and fetch the signed certificate
kubectl certificate approve ${csrName}
# verify certificate has been signed
for x in $(seq 20); do
    serverCert=$(kubectl get csr ${csrName} -o jsonpath='{.status.certificate}')
    if [[ ${serverCert} != '' ]]; then
        break
    fi
    sleep 1
done
if [[ ${serverCert} == '' ]]; then
    echo "ERROR: After approving csr ${csrName}, the signed certificate did not appear on the resource. Giving up after 20 attempts." >&2
    exit 1
fi
echo ${serverCert} | openssl base64 -d -A -out ${CERTDIR}/server-cert.pem

# ca cert
kubectl get configmap -n kube-system extension-apiserver-authentication -o=jsonpath='{.data.client-ca-file}' > ${CERTDIR}/ca-cert.pem

if [[ "${ENABLE_CREATE_SECRET}" = true ]]; then
    kubectl get ns ${NAMESPACE}
    if [ "$?" -eq 0 ]; then
        kubectl create ns ${NAMESPACE}
    fi
    # create the secret with CA cert and server cert/key
    kubectl create secret generic ${SECRET} \
        --from-file=tls.key=${CERTDIR}/server-key.pem \
        --from-file=tls.crt=${CERTDIR}/server-cert.pem \
        --from-file=ca.crt=${CERTDIR}/ca-cert.pem \
        --dry-run -o yaml |
    kubectl -n ${NAMESPACE} apply -f -
fi
