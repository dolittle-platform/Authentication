package viper

import (
	"fmt"
	"net/url"

	"dolittle.io/login/providers"
	"github.com/spf13/viper"
)

const (
	providersKey           = "providers"
	providerMapNameKey     = "name"
	providerMapImageURLKey = "image_url"
)

var (
	defaultProviders        = map[providers.ProviderID]providers.ProviderConfiguration(make(map[providers.ProviderID]providers.ProviderConfiguration, 0))
	defaultProviderName     = "UNKNOWN PROVIDER NAME"
	defaultProviderImageURL = "file:///etc/login/config/providers/images/default.png"
)

type providersConfiguration struct{}

func (c *providersConfiguration) Providers() map[providers.ProviderID]providers.ProviderConfiguration {
	if value := viper.GetStringMap(providersKey); value != nil && len(value) > 0 {
		return getProviderConfigurationMap(value)
	}
	return defaultProviders
}

func getProviderConfigurationMap(in map[string]interface{}) map[providers.ProviderID]providers.ProviderConfiguration {
	providerConfigMap := make(map[providers.ProviderID]providers.ProviderConfiguration, len(in))

	return providerConfigMap
}

func getProviderConfiguration(providerID providers.ProviderID, in interface{}) (providers.ProviderConfiguration, error) {
	name, ok := getProviderMapValue(providerID, providerMapNameKey).(string)

	if !ok {
		name = defaultProviderName
	}

	imageURLString, ok := getProviderMapValue(providerID, providerMapNameKey).(string)

	if !ok {
		imageURLString = defaultProviderImageURL
	}

	imageURL, err := url.Parse(imageURLString)
	if err != nil {
		return providers.ProviderConfiguration{}, err
	}

	return providers.ProviderConfiguration{
		Name:     name,
		ImageURL: imageURL,
	}, nil
}

func getProviderMapValue(providerID providers.ProviderID, key string) interface{} {
	return viper.Get(fmt.Sprint("%s.%s.%s", providersKey, string(providerID), key))
}
