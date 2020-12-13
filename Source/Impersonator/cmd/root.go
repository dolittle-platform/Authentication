package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var root = &cobra.Command{
	Use:   "impersonator",
	Short: `Impersonator is a Kubernetes API server proxy that impersonates users based on trusted HTTP headers`,
}

func init() {
	root.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Config file (default is $HOME/.impersonator.yaml")

	root.AddCommand(serve)
}

// Execute starts the program
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
