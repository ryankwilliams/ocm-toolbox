package filteredclusters

import (
	"github.com/ryankwilliams/ocm-toolbox/pkg/filteredclusters"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "filtered-clusters",
	Short: "Displays clusters filtered based on inputs",
	Long:  "Displays clusters filtered based on inputs",
	Run:   run,
}

func init() {}

func run(cmd *cobra.Command, argv []string) {
	filteredclusters.FilteredClusters(&filteredclusters.ClusterFilters{})
}
