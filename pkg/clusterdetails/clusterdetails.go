package clusterdetails

import (
	"fmt"

	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

func ClusterDetails() {
	ocm := ocm.Connect()

	clusters := ocm.ListClusters()

	if clusters.Len() == 0 {
		fmt.Printf("No clusters active in OCM %s", ocm.Connection.URL())
		return
	}

	for _, cluster := range clusters.Slice() {
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
  DELETION            : %s
  CLUSTER ACCESS      :
    $ ocm-toolbox cluster-credentials --cluster-id %s
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
		cluster.ExpirationTimestamp(),
		cluster.ID(),
		cluster.Name())
		fmt.Println()
	}
}
