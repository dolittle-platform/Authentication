package configuration

import (
	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/initiation"
	openid "dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/server"
	"dolittle.io/pascal/sessions"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Sessions() sessions.Configuration
	Initiation() initiation.Configuration
	Cookies() cookies.Configuration
	OpenID() openid.Configuration
}
