package options

import (
	flag "github.com/spf13/pflag"
)

// Config admission-controller server config.
type Config struct {
	Master                    string
	Kubeconfig                string
	CertFile                  string
	KeyFile                   string
	CaCertFile                string
	Port                      int
	MutateWebhookConfigName   string
	MutateWebhookName         string
	ValidateWebhookConfigName string
	ValidateWebhookName       string
	PrintVersion              bool
	AdmissionServiceName      string
	AdmissionServiceNamespace string
	SchedulerName             string
}

// NewConfig create new config
func NewConfig() *Config {
	c := Config{}
	return &c
}

// AddFlags add flags
func (c *Config) AddFlags(flag *flag.FlagSet) {
	flag.StringVar(&c.Master, "master", c.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	flag.StringVar(&c.Kubeconfig, "kubeconfig", c.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	flag.StringVar(&c.CertFile, "tls-cert-file", c.CertFile, ""+
		"File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
		"after server cert).")
	flag.StringVar(&c.KeyFile, "tls-private-key-file", c.KeyFile, "File containing the default x509 private key matching --tls-cert-file.")
	flag.StringVar(&c.CaCertFile, "ca-cert-file", c.CaCertFile, "File containing the x509 Certificate for HTTPS.")
	flag.IntVar(&c.Port, "port", 443, "the port used by admission-controller-server.")
	flag.StringVar(&c.MutateWebhookConfigName, "mutate-webhook-config-name", "",
		"Name of the mutatingwebhookconfiguration resource in Kubernetes [Deprecated]: it will be generated when not specified.")
	flag.StringVar(&c.MutateWebhookName, "mutate-webhook-name", "",
		"Name of the webhook entry in the webhook config. [Deprecated]: it will be generated when not specified")
	flag.StringVar(&c.ValidateWebhookConfigName, "validate-webhook-config-name", "",
		"Name of the mutatingwebhookconfiguration resource in Kubernetes. [Deprecated]: it will be generated when not specified")
	flag.StringVar(&c.ValidateWebhookName, "validate-webhook-name", "",
		"Name of the webhook entry in the webhook config. [Deprecated]: it will be generated when not specified")
	flag.BoolVar(&c.PrintVersion, "version", false, "Show version and quit")
	flag.StringVar(&c.AdmissionServiceNamespace, "webhook-namespace", "default", "The namespace of this webhook")
	flag.StringVar(&c.AdmissionServiceName, "webhook-service-name", "admission-service", "The name of this admission service")
}
