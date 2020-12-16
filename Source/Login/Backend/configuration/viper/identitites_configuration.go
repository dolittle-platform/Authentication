package viper

import (
	"dolittle.io/login/identities/tenants"
	"github.com/spf13/viper"
)

const (
	identitiesCurrentUserCookieNameKey = "identities.cookie_name"
	identitiesTenantNamesKey           = "identities.tenants"

	defaultIdentitiesCurrentUserCookieNameValue = "ory_kratos_session"
)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) Cookie() string {
	if value := viper.GetString(identitiesCurrentUserCookieNameKey); value != "" {
		return value
	}
	return defaultIdentitiesCurrentUserCookieNameValue
}

func (c *identitiesConfiguration) TenantNames() map[tenants.TenantID]string {
	if !viper.IsSet(identitiesTenantNamesKey) {
		return nil
	}
	tenantNames := map[tenants.TenantID]string{}
	for id, name := range viper.GetStringMapString(identitiesTenantNamesKey) {
		tenantNames[tenants.TenantID(id)] = name
	}
	return tenantNames
}
