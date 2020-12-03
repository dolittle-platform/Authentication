package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var root = &cobra.Command{
	Use: "pascal",
	Short: `Pascal is an OpenID Connect client that stores tokens in browser cookies, that can
be used as a middleware for web servers to enable OpenID Connect authentication.`,
}

func init() {
	root.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Config file (default is $HOME/.pascal.yaml)")

	root.AddCommand(serve)
}

// Execute starts the program
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
