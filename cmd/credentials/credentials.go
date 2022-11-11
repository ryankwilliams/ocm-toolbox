package credentials

import (
	"github.com/ryankwilliams/ocm-toolbox/pkg/credentials"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "cluster-credentials",
	Short: "Get credentials about a given cluster",
	Long: "Get credentials about a given cluster",
	Run: run,
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
		"OCM Cluster ID",
	)
	Cmd.MarkFlagRequired("cluster-id")
}

func run(cmd *cobra.Command, argv []string) {
	credentials.GetClusterCredentials(flags.clusterID)
}
