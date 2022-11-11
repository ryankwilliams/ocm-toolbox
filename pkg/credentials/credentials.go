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
	err := os.WriteFile(kubeconfigFilename, []byte(clusterCredentials.Kubeconfig()), 0600)

	if err != nil {
		fmt.Printf("Failed to write kubeconfig to disk, error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Cluster: %s, kubeconfig: %s\n", clusterName, kubeconfigFilename)

	adminPasswordFilename := clusterName + "-kubeadmin-password"
	err = os.WriteFile(adminPasswordFilename, []byte(clusterCredentials.Admin().Password()), 0600)

	if err != nil {
		fmt.Printf("Failed to write kubeadmin password to disk, error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Cluster: %s, kubeadmin password: %s\n", clusterName, adminPasswordFilename)
}
