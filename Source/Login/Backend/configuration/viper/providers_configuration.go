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

var (
	defaultProviderName     = "UNKNOWN PROVIDER"
	defaultProviderImageURL = ""
)

func newNilProvider() *providers.ProviderConfiguration {
	return &providers.ProviderConfiguration{
		Name:     defaultProviderName,
		ImageURL: nil,
	}
}

type providersConfiguration struct{}

func (c *providersConfiguration) Providers() providers.Providers {
	if !viper.IsSet(providersKey) {
		return nil
	}
	return getProviderConfigurationMap(viper.GetStringMap(providersKey))
}

func getProviderConfigurationMap(in map[string]interface{}) providers.Providers {
	providerConfigMap := make(map[providers.ProviderID]*providers.ProviderConfiguration, len(in))
	for provider, config := range in {
		providerConfigMap[provider] = getProviderConfiguration(provider, config)
	}
	return providers.Providers(providerConfigMap)
}

func getProviderConfiguration(providerID providers.ProviderID, config interface{}) *providers.ProviderConfiguration {
	configMap, ok := config.(map[string]interface{})
	if !ok {
		return newNilProvider()
	}
	name, ok := configMap[providersNameKey].(string)

	if !ok {
		return newNilProvider()
	}

	imageURLString, ok := configMap[providersImageURLKey].(string)

	if !ok {
		return newNilProvider()
	}

	imageURL, err := url.Parse(imageURLString)
	if err != nil {
		return newNilProvider()
	}

	return &providers.ProviderConfiguration{
		Name:     name,
		ImageURL: imageURL,
	}
}
