package viper

import (
	"strings"

	config "dolittle.io/impersonator/configuration"
	"dolittle.io/impersonator/identities"
	"dolittle.io/impersonator/proxy"
	"dolittle.io/impersonator/server"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func NewViperConfiguration(configPath string) (config.Configuration, error) {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".impersonator")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.WatchConfig()

	return &configuration{}, nil
}

type configuration struct {
	server     *serverConfiguration
	identities *identitiesConfiguration
	proxy      *proxyConfiguration
}

func (c *configuration) OnChange(callback func()) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		callback()
	})
}

func (c *configuration) Server() server.Configuration {
	return c.server
}

func (c *configuration) Identities() identities.Configuration {
	return c.identities
}

func (c *configuration) Proxy() proxy.Configuration {
	return c.proxy
}
