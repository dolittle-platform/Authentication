package flows

import (
	"dolittle.io/login/flows/consent"
	"dolittle.io/login/flows/login"
	"dolittle.io/login/flows/logout"
	"dolittle.io/login/flows/tenant"
)

type Configuration interface {
	Consent() consent.Configuration
	Login() login.Configuration
	Tenant() tenant.Configuration
	Logout() logout.Configuration
}
