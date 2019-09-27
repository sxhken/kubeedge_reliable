/*
Copyright 2019 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package device_crd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kubeedge/kubeedge/tests/e2e/constants"
	"github.com/kubeedge/kubeedge/tests/e2e/utils"
)

//context to load config and access across the package
var (
	ctx          *utils.TestContext
	nodeSelector string
	NodeName     string
)
var (
	deviceCRDPath      = "../../../build/crds/devices/devices_v1alpha1_device.yaml"
	deviceModelCRDPath = "../../../build/crds/devices/devices_v1alpha1_devicemodel.yaml"
	crdHandler         = "/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions"
	deviceCRD          = "devices.devices.kubeedge.io"
	deviceModelCRD     = "devicemodels.devices.kubeedge.io"
)

//Function to run the Ginkgo Test
func TestEdgecoreAppDeployment(t *testing.T) {

	RegisterFailHandler(Fail)
	var _ = BeforeSuite(func() {
		utils.Infof("Before Suite Execution")
		ctx = utils.NewTestContext(utils.LoadConfig())
		NodeName = "integration-node-" + utils.GetRandomString(10)
		nodeSelector = "node-" + utils.GetRandomString(3)

		//Check node successfully registered or not
		Eventually(func() string {
			status := utils.CheckNodeReadyStatus(ctx.Cfg.K8SMasterForKubeEdge+constants.NodeHandler, NodeName)
			utils.Infof("Node Name: %v, Node Status: %v", NodeName, status)
			return status
		}, "60s", "4s").Should(Equal("Running"), "Node register to the k8s master is unsuccessfull !!")

		err := utils.MqttConnect()
		Expect(err).To(BeNil())
	})
	AfterSuite(func() {
		By("After Suite Execution....!")
	})
	RunSpecs(t, "kubeedge Device Managemnet Suite")
}

func getpwd() string {
	_, file, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(file))
	Expect(err).Should(BeNil())
	return dir
}
