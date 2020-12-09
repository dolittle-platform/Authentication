package viper

import (
	"strings"

	"dolittle.io/login/clients"
	config "dolittle.io/login/configuration"
	clientsConfig "dolittle.io/login/configuration/viper/clients"
	flowsConfig "dolittle.io/login/configuration/viper/flows"
	"dolittle.io/login/flows"
	"dolittle.io/login/identities"
	"dolittle.io/login/server"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func NewViperConfiguration(configPath string, devServer bool) (config.Configuration, error) {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".login")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.WatchConfig()

	return &configuration{
		server: &serverConfiguration{
			devMode: devServer,
		},
		flows: &flowsConfiguration{
			consent: &flowsConfig.Consent{},
			login:   &flowsConfig.Login{},
			tenant:  &flowsConfig.Tenant{},
		},
		clients: &clientsConfiguration{
			hydra:  &clientsConfig.Hydra{},
			kratos: &clientsConfig.Kratos{},
		},
		identitites: &identitiesConfiguration{},
	}, nil
}

type configuration struct {
	server      *serverConfiguration
	flows       *flowsConfiguration
	clients     *clientsConfiguration
	identitites *identitiesConfiguration
}

func (c *configuration) OnChange(callback func()) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		callback()
	})
}

func (c *configuration) Server() server.Configuration {
	return c.server
}

func (c *configuration) Flows() flows.Configuration {
	return c.flows
}

func (c *configuration) Clients() clients.Configuration {
	return c.clients
}

func (c *configuration) Identities() identities.Configuration {
	return c.identitites
}
