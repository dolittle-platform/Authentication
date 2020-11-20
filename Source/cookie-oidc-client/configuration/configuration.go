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

type Configuration struct {
	Host                     string   `mapstructure:"host"`
	Port                     uint     `mapstructure:"port"`
	ProviderHost             string   `mapstructure:"provider_host"`
	ProviderPort             uint     `mapstructure:"provider_port"`
	ClientID                 string   `mapstructure:"client_id"`
	ClientSecret             string   `mapstructure:"client_secret"`
	RedirectBaseURL          string   `mapstructure:"redirect_base_url"`
	Scopes                   []string `mapstructure:"scopes"`
	SessionStoreName         string   `mapstructure:"session_store_name"`
	TokenCookieName          string   `mapstructure:"token_cookie_name"`
	TokenCookiePath          string   `mapstructure:"token_cookie_path"`
	TokenCookieExpiresInDays uint     `mapstructure:"token_cookie_expires_in_days"`
}

func GetDefaults() Configuration {
	return Configuration{
		Host:                     "http://localhost",
		Port:                     8888,
		ProviderHost:             "http://localhost",
		ProviderPort:             8080,
		ClientID:                 "do",
		ClientSecret:             "little",
		RedirectBaseURL:          "http://localhost:8080",
		Scopes:                   []string{oidc.ScopeOpenID},
		SessionStoreName:         "dolittle-session",
		TokenCookieName:          "dolittle-token",
		TokenCookiePath:          "/",
		TokenCookieExpiresInDays: 30,
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

func getConfigFilePath() string {
	return fmt.Sprintf("%s/%s", configPath, configName)
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
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	var configMap map[string]interface{}
	mapstructure.Decode(config, &configMap)
	log.Println(fmt.Sprintf("Configuration %v", configMap))

	return &config, nil
}
