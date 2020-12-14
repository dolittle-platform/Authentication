package providers

import "net/url"

type ProviderConfiguration struct {
	Name     string
	ImageURL *url.URL
}

type Providers = map[ProviderID]*ProviderConfiguration

type Configuration interface {
	Providers() Providers
}
