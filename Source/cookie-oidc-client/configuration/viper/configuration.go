package viper

import (
	"net/url"

	config "dolittle.io/cookie-oidc-client/configuration"
	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/server"
	"dolittle.io/cookie-oidc-client/sessions"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
)

func NewViperConfiguration() (config.Configuration, error) {
	return configuration{}, nil
}

type configuration struct {
	server     serverConfiguration
	initiation initiationConfiguration
	sessions   sessionsConfiguration
}

func (c configuration) Server() server.Configuration {
	return c.server
}

func (c configuration) Initiation() initiation.Configuration {
	return c.initiation
}

func (c configuration) Sessions() sessions.Configuration {
	return c.sessions
}

type serverConfiguration struct{}

func (c serverConfiguration) Port() int {
	return 8080
}

func (c serverConfiguration) InitiatePath() string {
	return "/initiate"
}

func (c serverConfiguration) CompletePath() string {
	return "/complete"
}

func (c serverConfiguration) ErrorRedirect() *url.URL {
	url, _ := url.Parse("http://localhost:8080/error")
	return url
}

type initiationConfiguration struct{}

func (initiationConfiguration) ReturnToParameter() string {
	return "return_to"
}

func (initiationConfiguration) DefaultReturnTo() *url.URL {
	url, _ := url.Parse("http://localhost:8080/return")
	return url
}

func (initiationConfiguration) AllowedReturnTo() []*url.URL {
	rl, _ := url.Parse("http://localhost:8080/return")
	return []*url.URL{rl}
}

type sessionsConfiguration struct {
	nonce nonceConfiguration
}

func (sessionsConfiguration) CookieName() string {
	return "dolittle"
}

func (s sessionsConfiguration) Nonce() nonces.Configuration {
	return s.nonce
}

type nonceConfiguration struct{}

func (nonceConfiguration) Size() int {
	return 18
}
