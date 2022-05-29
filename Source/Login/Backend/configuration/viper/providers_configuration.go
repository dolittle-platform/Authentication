package viper

import (
	"net/url"

	"dolittle.io/login/providers"
	"github.com/spf13/viper"
)

const (
	providersKey         = "providers"
	providersNameKey     = "name"
	providersImageURLKey = "image_url"
)

type providersConfiguration struct{}

func (c *providersConfiguration) Providers() providers.Providers {
	if !viper.IsSet(providersKey) {
		return nil
	}
	return getProviderConfigurationMap(viper.GetStringMap(providersKey))
}

func getProviderConfigurationMap(in map[string]interface{}) providers.Providers {
	providerConfigMap := make(map[providers.ProviderID]*providers.ProviderConfiguration)
	for providerID, config := range in {
		provider, ok := getProviderConfiguration(providerID, config)
		if ok {
			providerConfigMap[providerID] = provider
		}
	}
	return providers.Providers(providerConfigMap)
}

func getProviderConfiguration(providerID providers.ProviderID, config interface{}) (*providers.ProviderConfiguration, bool) {
	configMap, ok := config.(map[string]interface{})
	if !ok {
		return nil, false
	}
	name, ok := configMap[providersNameKey].(string)

	if !ok {
		return nil, false
	}

	imageURLString, ok := configMap[providersImageURLKey].(string)

	if !ok {
		return nil, false
	}

	imageURL, err := url.Parse(imageURLString)
	if err != nil {
		return nil, false
	}

	return &providers.ProviderConfiguration{
		Name:     name,
		ImageURL: imageURL,
	}, true
}
