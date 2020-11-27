package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello world")
		return nil
	},
}
