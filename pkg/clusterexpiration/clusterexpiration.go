package clusterexpiration

import (
	"fmt"
	"os"
	"time"

	"github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

func SetClusterExpirationTimestamp(clusterID string, duration int) {
	ocm := ocm.Connect()

	cluster, response := ocm.GetCluster(clusterID)
	clusterDetails := response.Body()

	fmt.Printf("Original expiration timestamp: %s\n", clusterDetails.ExpirationTimestamp())

	patchedCluster, err := v1.NewCluster().
		ExpirationTimestamp(clusterDetails.ExpirationTimestamp().Add(time.Duration(duration) * time.Minute)).
		Build()
	if err != nil {
		fmt.Printf("Unable to patch cluster expiration timestamp: %s, error: %s", clusterID, err)
		os.Exit(1)
	}

	_, err = cluster.Update().Body(patchedCluster).Send()
	if err != nil {
		fmt.Printf("Failed to update cluster expiration timestamp: %s, error: %s", clusterID, err)
		os.Exit(1)
	}

	fmt.Printf("Updated expiration timestamp: %s\n", patchedCluster.ExpirationTimestamp())
}
