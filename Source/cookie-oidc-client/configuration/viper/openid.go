package viper

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	openidIssuerKey       = "openid.issuer"
	openidClientIDKey     = "openid.client.id"
	openidClientSecretKey = "openid.client.secret"
	openidScopesKey       = "openid.scopes"
	openidRedirectURLKey  = "openid.redirect"
)

type openidConfiguration struct{}

func (c *openidConfiguration) Issuer() *url.URL {
	url, _ := url.Parse(viper.GetString(openidIssuerKey))
	return url
}

func (c *openidConfiguration) ClientID() string {
	return viper.GetString(openidClientIDKey)
}

func (c *openidConfiguration) ClientSecret() string {
	return viper.GetString(openidClientSecretKey)
}

func (c *openidConfiguration) Scopes() []string {
	return viper.GetStringSlice(openidScopesKey)
}

func (c *openidConfiguration) RedirectURL() *url.URL {
	url, _ := url.Parse(viper.GetString(openidRedirectURLKey))
	return url
}
