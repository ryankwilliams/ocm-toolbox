package clusterdetails

import (
	"github.com/ryankwilliams/ocm-toolbox/pkg/clusterdetails"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cluster-details",
	Short: "Get cluster details beyond the defaults",
	Long:  "Get cluster details beyond the defaults",
	Run:   run,
}

func init() {}

func run(cmd *cobra.Command, argv []string) {
	clusterdetails.ClusterDetails()
}
