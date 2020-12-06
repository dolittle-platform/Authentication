package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "login",
	Short: `TODO.`,
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Config file (default is $HOME/.login.yaml)")
	rootCmd.AddCommand(serveCmd)
}

// Execute starts the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
