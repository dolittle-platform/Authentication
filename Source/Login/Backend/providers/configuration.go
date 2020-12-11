package providers

import "net/url"

type ProviderConfiguration struct {
	Name     string
	ImageURL *url.URL
}

type Configuration interface {
	Providers() map[ProviderID]ProviderConfiguration
}
