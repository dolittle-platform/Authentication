package configuration

import (
	"dolittle.io/login/clients"
	"dolittle.io/login/flows"
	"dolittle.io/login/identities"
	"dolittle.io/login/server"
)

type Configuration interface {
	OnChange(callback func())

	Server() server.Configuration
	Clients() clients.Configuration
	Flows() flows.Configuration
	Identities() identities.Configuration
}
