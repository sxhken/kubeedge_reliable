package app

import (
	"github.com/spf13/cobra"

	admissioncontroller "github.com/kubeedge/kubeedge/cloud/pkg/admissioncontroller/controller"
)

func NewAdmissionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "admission",
		Long: `Admission leverage the feature of Dynamic Admission Control from kubernetes, start it
if want to admission control some kubeedge resources.`,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}

	return cmd
}
