package viper

import (
	"strings"

	config "dolittle.io/cookie-oidc-client/configuration"
	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/server"
	"dolittle.io/cookie-oidc-client/sessions"
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
		viper.SetConfigName(".cookie-oidc-client")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.WatchConfig()

	return &configuration{
		server:     &serverConfiguration{},
		initiation: &initiationConfiguration{},
		sessions: &sessionsConfiguration{
			nonce: &nonceConfiguration{},
		},
	}, nil
}

type configuration struct {
	server     *serverConfiguration
	initiation *initiationConfiguration
	sessions   *sessionsConfiguration
}

func (c *configuration) OnChange(callback func()) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		callback()
	})
}

func (c *configuration) Server() server.Configuration {
	return c.server
}

func (c *configuration) Initiation() initiation.Configuration {
	return c.initiation
}

func (c *configuration) Sessions() sessions.Configuration {
	return c.sessions
}
