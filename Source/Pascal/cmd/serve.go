package cmd

import (
	"dolittle.io/pascal/configuration"
	"dolittle.io/pascal/configuration/viper"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Pascal server",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := viper.NewViperConfiguration(configPath)
		if err != nil {
			return err
		}
		container, err := configuration.NewContainer(config)
		if err != nil {
			return err
		}
		return container.Server.Run()
	},
}
