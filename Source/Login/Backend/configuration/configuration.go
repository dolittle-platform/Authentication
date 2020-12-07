package configuration

import (
	"dolittle.io/login/clients"
	"dolittle.io/login/flows"
	"dolittle.io/login/identities"
	"dolittle.io/login/server"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Clients() clients.Configuration
	Flows() flows.Configuration
	Identities() identities.Configuration
}

// package configuration

// import (
// 	"fmt"
// 	"log"

// 	"github.com/mitchellh/mapstructure"
// 	"github.com/spf13/viper"
// )

// const (
// 	configName = "config"
// 	configType = "yaml"
// 	configPath = "/etc/tenant-selector"
// )

// type Configuration struct {
// 	Host            string `mapstructure:"host"`
// 	Port            uint   `mapstructure:"port"`
// 	HydraAdminURL   string `mapstructure:"hydra_admin_url"`
// 	KratosPublicURL string `mapstructure:"kratos_public_url"`
// }

// func GetDefaults() Configuration {
// 	return Configuration{
// 		Host:            "http://localhost",
// 		Port:            8889,
// 		HydraAdminURL:   "http://localhost:4445",
// 		KratosPublicURL: "http://localhost:8080/.ory/kratos/public/",
// 	}
// }

// func Setup(defaultConfig *Configuration) error {
// 	var defaultsMap map[string]interface{}
// 	err := mapstructure.Decode(defaultConfig, &defaultsMap)
// 	if err != nil {
// 		return err
// 	}
// 	for k, v := range defaultsMap {
// 		viper.SetDefault(k, v)
// 	}
// 	viper.SetConfigName(configName)
// 	viper.AddConfigPath(configPath)
// 	return nil
// }

// func Read() (*Configuration, error) {
// 	log.Println(fmt.Sprintf("Reading configuration from '%s'", configPath))
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Println(fmt.Sprintf("Could not find configuration file '%s' at '%s'. Using default values", configName, configPath))
// 	} else {
// 		log.Println(fmt.Sprintf("Used configuration'%s'", viper.ConfigFileUsed()))
// 	}
// 	config := Configuration{}
// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var configMap map[string]interface{}
// 	mapstructure.Decode(config, &configMap)
// 	log.Println(fmt.Sprintf("Configuration %v", configMap))

// 	return &config, nil
// }
