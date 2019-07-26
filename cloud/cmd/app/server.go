package app

import (
	"github.com/spf13/cobra"

	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/kubeedge/cloud/pkg/cloudhub"
	"github.com/kubeedge/kubeedge/cloud/pkg/controller"
	"github.com/kubeedge/kubeedge/cloud/pkg/devicecontroller"
)

func NewEdgeControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "EdgeController",
		Long: `The Kubernetes scheduler is a policy-rich, topology-aware,
workload-specific function that significantly impacts availability, performance,
and capacity. The scheduler needs to take into account individual and collective
resource requirements, quality of service requirements, hardware/software/policy
constraints, affinity and anti-affinity specifications, data locality, inter-workload
interference, deadlines, and so on. Workload-specific requirements will be exposed
through the API as necessary.`,
		Run: func(cmd *cobra.Command, args []string) {
			registerModules()
			// start all modules
			core.Run()
		},
	}

	return cmd
}

// registerModules register all the modules started in edgecontroller
func registerModules() {
	cloudhub.Register()
	controller.Register()
	devicecontroller.Register()
}
