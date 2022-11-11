package main

import (
	"os"

	"github.com/ryankwilliams/ocm-toolbox/cmd/credentials"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ocm-toolbox",
	Short: "Helpful commands used on a daily basis while working with OCM",
	Long: "Helpful commands used on a daily basis while working with OCM",
	Run: run,
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var flags struct {
	token string
	url string
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&flags.token,
		"token",
		"",
		"",
		"OCM API Token",
	)

	rootCmd.PersistentFlags().StringVarP(
		&flags.url,
		"url",
		"",
		"https://api.openshift.com",
		"OCM API URL",
	)

	viper.SetDefault("ocmToken", os.Getenv("OCM_TOKEN"))
	viper.BindPFlag("ocmToken", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("ocmUrl", rootCmd.PersistentFlags().Lookup("url"))

	rootCmd.AddCommand(credentials.Cmd)
}

func run(cmd *cobra.Command, argv []string) {}
