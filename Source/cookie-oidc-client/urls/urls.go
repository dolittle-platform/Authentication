package urls

import (
	"fmt"
	"net/url"

	"dolittle.io/cookie-oidc-client/configuration"
)

type OIDCProviderIssuerGetter interface {
	GetOIDCProviderIssuerURL() (*url.URL, error)
}

type oIDCProviderIssuerGetter struct {
	config *configuration.Configuration
}

var _ OIDCProviderIssuerGetter = new(oIDCProviderIssuerGetter)

func NewOIDCProviderIssuerGetter(config *configuration.Configuration) *oIDCProviderIssuerGetter {
	return &oIDCProviderIssuerGetter{config}
}

func (self *oIDCProviderIssuerGetter) GetOIDCProviderIssuerURL() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s:%d/", self.config.Provider.Host, self.config.Provider.Port))
}
