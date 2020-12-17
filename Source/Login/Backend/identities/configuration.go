package identities

import (
	"dolittle.io/login/identities/current"
	"dolittle.io/login/identities/tenants"
)

type Configuration interface {
	current.Configuration
	tenants.Configuration
}
