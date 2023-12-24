package ocm

import (
	"context"
	"fmt"
	"os"
	"regexp"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
	"github.com/spf13/viper"
)

type OCMInstance struct {
	Connection *sdk.Connection
}

func GetOcmApiUrl() (string, string) {
	ocmUrl := viper.GetString("ocmUrl")
	ocmShort := ""

	switch ocmUrl {
	case "https://api.openshift.com", "prod":
		ocmUrl = "https://api.openshift.com"
		ocmShort = "prod"
	case "https://api.stage.openshift.com", "staging":
		ocmUrl = "https://api.stage.openshift.com"
		ocmShort = "staging"
	case "https://api.integration.openshift.com", "integration":
		ocmUrl = "https://api.integration.openshift.com"
		ocmShort = "integration"
	}

	return ocmUrl, ocmShort
}

func Connect() *OCMInstance {
	token := viper.GetString("ocmToken")

	ctx := context.Background()

	logger, err := logging.NewGoLoggerBuilder().
		Debug(false).
		Build()

	if err != nil {
		panic(err)
	}

	apiUrl, _ := GetOcmApiUrl()

	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(apiUrl).
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

func (o *OCMInstance) GetCluster(clusterID string) (*v1.ClusterClient, *v1.ClusterGetResponse) {
	cluster := o.Connection.ClustersMgmt().V1().Clusters().Cluster(clusterID)

	response, err := cluster.Get().Send()
	if err != nil {
		fmt.Printf("Unable to find cluster: %s in OCM, error: %s", clusterID, err)
		os.Exit(1)
	}

	return cluster, response
}

func (o *OCMInstance) GetClusterBody(clusterID string) *v1.Cluster {
	_, response := o.GetCluster(clusterID)
	return response.Body()
}

func (o *OCMInstance) ListClustersRegex(regex string) []*v1.Cluster {
	clusters := o.ListClusters()

	var clusterList []*v1.Cluster

	for _, cluster := range clusters.Slice() {
		if match, _ := regexp.MatchString(regex, cluster.Name()); match {
			clusterList = append(clusterList, cluster)
		}
	}
	return clusterList
}

func (o *OCMInstance) ListManagedClusters(clusters []*v1.Cluster) []*v1.Cluster {
	var clusterList []*v1.Cluster

	for _, cluster := range clusters {
	    cluster_id := cluster.Product().ID()
		if cluster_type != "ocp" && cluster_type != "aro" {
			clusterList = append(clusterList, cluster)
		}
	}
	return clusterList
}

func (o *OCMInstance) GetClusterCredentials(clusterID string) (string, *v1.ClusterCredentials) {
	cluster := o.GetClusterBody(clusterID)

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
