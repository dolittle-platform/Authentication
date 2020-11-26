package handlers

import (
	"net/http"
	"net/url"

	"dolittle.io/cookie-oidc-client/configuration"
	"golang.org/x/oauth2"
)

type CallbackRedirectGetter interface {
	GetCallbackRedirectURL() (*url.URL, error)
}

type callbackRedirectGetter struct {
	config *configuration.Configuration
}

var _ CallbackRedirectGetter = new(callbackRedirectGetter)

func NewCallbackRedirectGetter(config *configuration.Configuration) *callbackRedirectGetter {
	return &callbackRedirectGetter{config}
}

func (self *callbackRedirectGetter) GetCallbackRedirectURL() (*url.URL, error) {
	return url.Parse(self.config.CallbackRedirectURL)
}

type ReturnToGetter interface {
	GetReturnToURL(r *http.Request) (*url.URL, error)
}

type returnToGetter struct {
	defaultURL *url.URL
}

var _ ReturnToGetter = new(returnToGetter)

func NewReturnToGetter(config *configuration.Configuration) (*returnToGetter, error) {
	url, err := url.Parse(config.DefaultReturnTo)
	if err != nil {
		return nil, err
	}
	return &returnToGetter{url}, nil
}

func (self *returnToGetter) GetReturnToURL(r *http.Request) (*url.URL, error) {
	return url.Parse(r.URL.Query().Get("return_to"))
}

type ConsentPageGetter interface {
	GetConsentPageURL(nonce string) (*url.URL, error)
}

type consentPageGetter struct {
	oauthConfig oauth2.Config
}

var _ ConsentPageGetter = new(consentPageGetter)

func NewConsentPageGetter(config oauth2.Config) *consentPageGetter {
	return &consentPageGetter{config}
}

func (self *consentPageGetter) GetConsentPageURL(nonce string) (*url.URL, error) {
	return url.Parse(self.oauthConfig.AuthCodeURL(nonce))
}
