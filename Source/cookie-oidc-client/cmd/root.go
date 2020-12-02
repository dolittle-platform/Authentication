package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var root = &cobra.Command{
	Use:   "cookie-oidc-client",
	Short: "This is very cool",
}

func init() {
	root.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Config file (default is $HOME/.cookie-oidc-client.yaml)")

	root.AddCommand(serve)
}

// Execute starts the program
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
