package ocm

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
	"github.com/spf13/viper"
)

type OCMInstance struct {
	Connection *sdk.Connection
}

func Connect() *OCMInstance {
	token := viper.GetString("ocmToken")
	ocmUrl := viper.GetString("ocmUrl")

	ctx := context.Background()

	logger, err := logging.NewGoLoggerBuilder().
		Debug(false).
		Build()

	if err != nil {
		panic(err)
	}

	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(ocmUrl).
		BuildContext(ctx)

	if err != nil {
		panic(err)
	}

	return &OCMInstance{
		Connection: connection,
	}
}

func (o *OCMInstance) ListClusters() *v1.ClusterList {
	collection := o.Connection.ClustersMgmt().V1().Clusters().List()
	response, err := collection.Send()

	if err != nil {
		panic(err)
	}

	clusters := response.Items()

	return clusters
}

func (o *OCMInstance) GetCluster(clusterID string) *v1.Cluster {
	response, err := o.Connection.ClustersMgmt().V1().Clusters().Cluster(clusterID).Get().Send()

	if err != nil {
		fmt.Printf("Unable to find cluster: %s in OCM, error: %s", clusterID, err)
		os.Exit(1)
	}

	return response.Body()
}

func (o *OCMInstance) GetClusterCredentials(clusterID string) (string, *v1.ClusterCredentials) {
	cluster := o.GetCluster(clusterID)

	response, err := o.Connection.ClustersMgmt().V1().Clusters().
		Cluster(clusterID).
		Credentials().
		Get().
		Send()

	if err != nil {
		fmt.Printf("Unable to retrieve credentials for cluster: %s, error: %s", clusterID, err)
		os.Exit(1)
	}

	return cluster.Name(), response.Body()
}
