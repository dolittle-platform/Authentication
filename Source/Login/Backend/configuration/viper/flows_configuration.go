package viper

import (
	"dolittle.io/login/configuration/viper/flows"
	"dolittle.io/login/flows/consent"
	"dolittle.io/login/flows/login"
	"dolittle.io/login/flows/logout"
	"dolittle.io/login/flows/tenant"
)

type flowsConfiguration struct {
	consent *flows.Consent
	login   *flows.Login
	tenant  *flows.Tenant
	logout  *flows.Logout
}

func (c *flowsConfiguration) Consent() consent.Configuration {
	return c.consent
}

func (c *flowsConfiguration) Login() login.Configuration {
	return c.login
}

func (c *flowsConfiguration) Tenant() tenant.Configuration {
	return c.tenant
}

func (c *flowsConfiguration) Logout() logout.Configuration {
	return c.logout
}
