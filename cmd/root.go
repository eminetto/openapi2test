package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openapi2test",
	Short: "Generate Go tests from an OpenAPI specification",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	// add new commands here
}
