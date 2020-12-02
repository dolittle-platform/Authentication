package configuration

import (
	"dolittle.io/cookie-oidc-client/cookies"
	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/openid"
	"dolittle.io/cookie-oidc-client/server"
	"dolittle.io/cookie-oidc-client/sessions"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Sessions() sessions.Configuration
	Initiation() initiation.Configuration
	Cookies() cookies.Configuration
	OpenID() openid.Configuration
}
