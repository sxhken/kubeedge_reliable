package main

import (
	"os"

	"k8s.io/klog"

	"github.com/kubeedge/kubeedge/cloud/cmd/app"
)

func main() {
	command := app.NewEdgeControllerCommand()

	klog.InitFlags(nil)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
