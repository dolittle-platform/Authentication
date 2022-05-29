package configuration

import (
	"dolittle.io/pascal/cookies"
	openid "dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/redirects"
	"dolittle.io/pascal/server"
	"dolittle.io/pascal/sessions"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Sessions() sessions.Configuration
	Redirects() redirects.Configuration
	Cookies() cookies.Configuration
	OpenID() openid.Configuration
}
