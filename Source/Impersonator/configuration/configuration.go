package configuration

import (
	"dolittle.io/impersonator/identities"
	"dolittle.io/impersonator/proxy"
	"dolittle.io/impersonator/server"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Identities() identities.Configuration
	Proxy() proxy.Configuration
}
