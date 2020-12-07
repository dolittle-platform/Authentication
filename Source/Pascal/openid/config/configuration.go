package config

import "net/url"

type Configuration interface {
	Issuer() *url.URL
	ClientID() string
	ClientSecret() string
	Scopes() []string
	RedirectURL() *url.URL
}
