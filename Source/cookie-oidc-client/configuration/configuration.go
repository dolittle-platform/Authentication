package configuration

import (
	"fmt"
	"log"

	"github.com/coreos/go-oidc"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	configPath = "/etc/cookie-oidc-client"
)

/*
example yaml
host: hoster.io
port: 2020
provider:
	host: hydrahost
	port: 1800-hydra-ur-port-8080
	client_id: id
	client_secret: secreeeet
	redirect_url: http://localhost:8080/.auth/callback
	scopes:
		- id: lol
		- id: løl
cookie:
	name: dolittle-token
	path: /
	expires_in_days: 30
session_store_name: sneaky-store
*/

type provider struct {
	Host         string   `mapstructure:"host"`
	Port         uint     `mapstructure:"port"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	RedirectURL  string   `mapstructure:"redirect_url"`
	Scopes       []string `mapstructure:"scopes"`
}

type cookie struct {
	Name          string `mapstructure:"name"`
	Path          string `mapstructure:"path"`
	ExpiresInDays uint   `mapstructure:"expires_in_days"`
}

type Configuration struct {
	Host                string   `mapstructure:"host"`
	Port                uint     `mapstructure:"port"`
	Provider            provider `mapstructure:"provider"`
	Cookie              cookie   `mapstructure:"cookie"`
	SessionStoreName    string   `mapstructure:"session_store_name"`
	DefaultReturnTo     string   `mapstructure:"default_return_to"`
	CallbackRedirectURL string   `mapstructure:"callback_redirect_url"`
}

func GetDefaults() Configuration {
	return Configuration{
		Host: "http://localhost",
		Port: 8888,
		Provider: provider{
			Host:         "http://localhost",
			Port:         8080,
			ClientID:     "do",
			ClientSecret: "little",
			RedirectURL:  "http://localhost:8080/.auth/callback/",
			Scopes:       []string{oidc.ScopeOpenID},
		},
		Cookie: cookie{
			Name:          "dolittle-token",
			Path:          "/",
			ExpiresInDays: 30,
		},
		SessionStoreName:    "dolittle-session",
		DefaultReturnTo:     "http://localhost:8080/",
		CallbackRedirectURL: "http://localhost:8080/",
	}
}

func Setup(defaultConfig *Configuration) error {
	var defaultsMap map[string]interface{}
	err := mapstructure.Decode(defaultConfig, &defaultsMap)
	if err != nil {
		return err
	}
	for k, v := range defaultsMap {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	return nil
}

func Read() (*Configuration, error) {
	log.Println(fmt.Sprintf("Reading configuration from '%s'", configPath))
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(fmt.Sprintf("Could not find configuration file '%s' at '%s'. Using default values", configName, configPath))
	} else {
		log.Println(fmt.Sprintf("Used configuration'%s'", viper.ConfigFileUsed()))
	}
	config := Configuration{}
	return &config, viper.Unmarshal(&config)
}
