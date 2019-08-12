package main

import (
	"os"

	"github.com/kubeedge/kubeedge/cloud/cmd/admission/app"
)

func main() {
	command := app.NewAdmissionCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
