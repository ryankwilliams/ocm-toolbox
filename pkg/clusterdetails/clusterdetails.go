package clusterdetails

import (
	"fmt"
	"os"
	"time"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

type ClusterFilters struct {
	ID        string
	NameRegex string
}

var currentDirectory *string

func ClusterDetails(clusterFilters *ClusterFilters) {
	directory, err := os.Getwd()
	currentDirectory = &directory
	if err != nil {
		fmt.Printf("Unable to locate current working directory: %v", err)
		return
	}

	ocmInstance := ocm.Connect()

	var presentClusters []*v1.Cluster

	if clusterFilters.ID != "" {
		cluster := ocmInstance.GetClusterBody(clusterFilters.ID)
		presentClusters = make([]*v1.Cluster, 0)
		presentClusters = append(presentClusters, cluster)
	} else if clusterFilters.NameRegex != "" {
		presentClusters = ocmInstance.ListClustersRegex(clusterFilters.NameRegex)
	} else {
		clusters := ocmInstance.ListClusters()
		presentClusters = clusters.Slice()
	}

	// Prompt only OSD/ROSA clusters
	var filteredClusters []*v1.Cluster

	for _, cluster := range presentClusters {
		product_id := cluster.Product().ID()
		if product_id != "ocp" {
			filteredClusters = append(filteredClusters, cluster)
		}
	}

	if len(filteredClusters) == 0 {
		fmt.Printf("No clusters found in ocm %s\n", ocmInstance.Connection.URL())
		return
	}

	for _, cluster := range filteredClusters {
		creation := cluster.CreationTimestamp()
		creationTime := time.Date(
			creation.Year(),
			creation.Month(),
			creation.Day(),
			creation.Hour(),
			creation.Minute(),
			creation.Second(),
			creation.Nanosecond(),
			creation.Location())
		currentTime := time.Now().UTC()
		timeDiff := time.Time{}.Add(currentTime.Sub(creationTime))
		_, apiUrlShort := ocm.GetOcmApiUrl()

		fmt.Println(clusterInfoFormatted(*cluster, timeDiff, apiUrlShort))
	}

	fmt.Printf("Total clusters: %v\n", len(filteredClusters))
}

func clusterInfoFormatted(cluster v1.Cluster, timeDiff time.Time, apiUrlShort string) string {
	clusterInfo := fmt.Sprintf(`Cluster: %s
  ID                     : %s
  API URL                : %s
  CONSOLE URL            : %s
  OPENSHIFT VERSION      : %s
  PRODUCT ID             : %v`,
		cluster.Name(),
		cluster.ID(),
		cluster.API().URL(),
		cluster.Console().URL(),
		cluster.OpenshiftVersion(),
		cluster.Product().ID())

	if cluster.Product().ID() == "rosa" {
		clusterInfo += fmt.Sprintf("\n  ROSA HCP               : %v", cluster.Hypershift().Enabled())
	}

	clusterInfo += fmt.Sprintf(`
  CLOUD PROVIDER         : %s
  REGION                 : %s
  STATE                  : %s
  CONTROL PLANE NODES    : %v
  COMPUTE NODES          : %v
  CREATION               : %s
  UP TIME (H:M:S)        : %s
  DELETION               : %s
  CLUSTER ACCESS         :
    $ ocm-toolbox cluster-credentials --cluster-id %s --url %s
    $ export KUBECONFIG=%s/%s-kubeconfig
    $ oc cluster-info
  SET CLUSTER EXPIRATION :
    $ ocm-toolbox set-cluster-expiration --cluster-id %s --duration <duration-in-minutes> --url %s`,
		cluster.CloudProvider().ID(),
		cluster.Region().ID(),
		cluster.State(),
		cluster.Nodes().Master(),
		cluster.Nodes().Compute(),
		cluster.CreationTimestamp(),
		timeDiff.Format("15:4:5"),
		cluster.ExpirationTimestamp(),
		cluster.ID(),
		apiUrlShort,
		*currentDirectory,
		cluster.Name(),
		cluster.ID(),
		apiUrlShort,
	)

	return clusterInfo
}
