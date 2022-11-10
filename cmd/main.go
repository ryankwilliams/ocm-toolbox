package main

import (
	"os"

	"github.com/spf13/cobra"
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

func init() {}

func run(cmd *cobra.Command, argv []string) {}
