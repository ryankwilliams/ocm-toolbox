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

var flags struct {
	clusterID string
}

func init() {
	Cmd.Flags().StringVarP(
		&flags.clusterID,
		"cluster-id",
		"",
		"",
		"OCM Cluster ID (query for an exact match)",
	)
}

func run(cmd *cobra.Command, argv []string) {
	clusterdetails.ClusterDetails(&clusterdetails.ClusterFilters{
		ID: flags.clusterID,
	})
}
