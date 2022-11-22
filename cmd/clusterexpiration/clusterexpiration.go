package clusterexpiration

import (
	"github.com/ryankwilliams/ocm-toolbox/pkg/clusterexpiration"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "set-cluster-expiration",
	Short: "Update a clusters expiration date",
	Long:  "Update a clusters expiration date",
	Run:   run,
}

var flags struct {
	clusterID string
	duration  int
}

func init() {
	Cmd.Flags().StringVarP(
		&flags.clusterID,
		"cluster-id",
		"",
		"",
		"OCM Cluster ID",
	)
	Cmd.Flags().IntVarP(
		&flags.duration,
		"duration",
		"",
		60,
		"Total time (in minutes) to extend cluster expiration",
	)
	Cmd.MarkFlagRequired("cluster-id") // nolint:errcheck
}

func run(cmd *cobra.Command, argv []string) {
	clusterexpiration.SetClusterExpirationTimestamp(
		flags.clusterID,
		flags.duration,
	)
}
