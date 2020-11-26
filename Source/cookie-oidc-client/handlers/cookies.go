package handlers

import (
	"net/http"
	"time"

	"dolittle.io/cookie-oidc-client/configuration"
)

type CookieFactory interface {
	Create(value string) *http.Cookie
}

type oauth2CookieFactory struct {
	config *configuration.Configuration
}

func NewOauth2CookieFactory(config *configuration.Configuration) *oauth2CookieFactory {
	return &oauth2CookieFactory{config}
}

func (self *oauth2CookieFactory) Create(value string) *http.Cookie {
	return &http.Cookie{
		Name:    self.config.Cookie.Name,
		Value:   value,
		Path:    self.config.Cookie.Path,
		Expires: time.Now().Add(time.Duration(self.config.Cookie.ExpiresInDays) * 24 * time.Hour),
	}
}
