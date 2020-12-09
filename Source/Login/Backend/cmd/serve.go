package cmd

import (
	"dolittle.io/login/configuration"
	"dolittle.io/login/configuration/viper"
	"github.com/spf13/cobra"
)

var devServer bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Login server",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := viper.NewViperConfiguration(configPath, devServer)
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

func init() {
	serveCmd.Flags().BoolVar(&devServer, "dev", false, "Starts the server in DEV mode")
}
