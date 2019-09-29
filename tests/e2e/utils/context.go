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
package utils

import (
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//Test context struct
type TestContext struct {
	Cfg Config
}

//NewTestContext function to build testcontext with provided config.
func NewTestContext(cfg Config) *TestContext {
	return &TestContext{
		Cfg: cfg,
	}
}

//SendHttpRequest Function to prepare the http req and send
func SendHttpRequest(method, kubeconfig, requestURI string) (*rest.Result, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset := kubernetes.NewForConfigOrDie(config)

	result := rest.Result{}
	if method == http.MethodGet {
		result = clientset.RESTClient().Get().RequestURI(requestURI).Do()
	} else if method == http.MethodDelete {
		result = clientset.RESTClient().Delete().RequestURI(requestURI).Do()
	}

	if result.Error() != nil {
		// handle error
		Fatalf("HTTP request is failed :%v", result.Error())
		return nil, result.Error()
	}

	var statusCode *int
	result.StatusCode(statusCode)

	Infof("HTTP request is successful: %s %s %v", method, requestURI, statusCode)
	return &result, nil
}

//MapLabels function add label selector
func MapLabels(ls map[string]string) string {
	selector := make([]string, 0, len(ls))
	for key, value := range ls {
		selector = append(selector, key+"="+value)
	}
	sort.StringSlice(selector).Sort()
	return url.QueryEscape(strings.Join(selector, ","))
}
