package filteredclusters

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

type ClusterFilters struct{}

type ClusterSpec struct {
	Name      string          `json:"name"`
	CreatedOn string          `json:"createdOn"`
	Creator   string          `json:"creator"`
	State     v1.ClusterState `json:"state"`
}

type ClustersFiltered struct {
	Clusters []ClusterSpec `json:"clusters"`
}

func FilteredClusters(clusterFilters *ClusterFilters) {
	ocmInstance := ocm.Connect()
	clusters := ocmInstance.ListClusters()

	var filteredClusters = ""
	clustersFiltered := &ClustersFiltered{}

	for _, cluster := range clusters.Slice() {
		clusterSpec := ClusterSpec{
			Name:      cluster.Name(),
			CreatedOn: cluster.CreationTimestamp().Format(time.RFC822),
			Creator:   ocmInstance.GetClusterCreator(cluster),
			State:     cluster.State(),
		}

		filteredClusters += fmt.Sprintf("Name: '%s', Created On: '%s', Creator: '%s' "+
			"State: '%s'\n",
			clusterSpec.Name,
			clusterSpec.CreatedOn,
			clusterSpec.CreatedOn,
			clusterSpec.State,
		)

		clustersFiltered.Clusters = append(clustersFiltered.Clusters, clusterSpec)
	}

	data, _ := json.Marshal(clustersFiltered)
	os.WriteFile("filtered-clusters.json", data, 0644)

	fmt.Print(filteredClusters)
}
