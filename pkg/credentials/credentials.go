package credentials

import (
	"fmt"
	"os"

	"github.com/ryankwilliams/ocm-toolbox/pkg/ocm"
)

func GetClusterCredentials(clusterID string) {
	ocm := ocm.Connect()
	clusterName, clusterCredentials := ocm.GetClusterCredentials(clusterID)

	kubeconfigFilename := clusterName + "-kubeconfig"
	os.WriteFile(kubeconfigFilename, []byte(clusterCredentials.Kubeconfig()), 0600)
	fmt.Printf("Cluster: %s, kubeconfig: %s\n", clusterName, kubeconfigFilename)

	adminPasswordFilename := clusterName + "-kubeadmin-password"
	os.WriteFile(adminPasswordFilename, []byte(clusterCredentials.Admin().Password()), 0600)
	fmt.Printf("Cluster: %s, kubeadmin password: %s\n", clusterName, adminPasswordFilename)
}
