package clusterdetails

import (
	"fmt"
	"time"

	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

func ClusterDetails() {
	ocm := ocm.Connect()

	clusters := ocm.ListClusters()

	if clusters.Len() == 0 {
		fmt.Printf("No clusters active in OCM %s\n", ocm.Connection.URL())
		return
	}

	for _, cluster := range clusters.Slice() {
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

		fmt.Printf(`Cluster: %s
  ID                  : %s
  API URL             : %s
  CONSOLE URL         : %s
  OPENSHIFT VERSION   : %s
  PRODUCT ID          : %s
  CLOUD PROVIDER      : %s
  REGION              : %s
  STATE               : %s
  CONTROL PLANE NODES : %v
  COMPUTE NODES       : %v
  CREATION            : %s
  UP TIME (H:M:S)     : %s
  DELETION            : %s
  CLUSTER ACCESS      :
    $ ocm-toolbox cluster-credentials --cluster-id %s --url %s
    $ export KUBECONFIG=%s-kubeconfig
    $ oc cluster-info`,
			cluster.Name(),
			cluster.ID(),
			cluster.API().URL(),
			cluster.Console().URL(),
			cluster.OpenshiftVersion(),
			cluster.Product().ID(),
			cluster.CloudProvider().ID(),
			cluster.Region().ID(),
			cluster.State(),
			cluster.Nodes().Master(),
			cluster.Nodes().Compute(),
			cluster.CreationTimestamp(),
			timeDiff.Format("15:4:5"),
			cluster.ExpirationTimestamp(),
			cluster.ID(),
			ocm.Connection.URL(),
			cluster.Name())
		fmt.Println()
	}
}