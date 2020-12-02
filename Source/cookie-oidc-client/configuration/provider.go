package configuration

import (
	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/server"
	"dolittle.io/cookie-oidc-client/sessions"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Initiation() initiation.Configuration
	Sessions() sessions.Configuration
}
