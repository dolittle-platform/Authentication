package viper

import (
	"dolittle.io/login/identities/tenants"
	"github.com/spf13/viper"
)

const (
	identitiesCurrentUserCookieNameKey = "identities.cookie_name"
	identitiesTenantNameMapKey         = "identities.tenants"

	defaultIdentitiesCurrentUserCookieNameValue = "ory_kratos_session"
)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) Cookie() string {
	if value := viper.GetString(identitiesCurrentUserCookieNameKey); value != "" {
		return value
	}
	return defaultIdentitiesCurrentUserCookieNameValue
}

func (c *identitiesConfiguration) TenantNamesMap() map[tenants.TenantID]string {
	tenantNameMapRaw := viper.GetStringMapString(identitiesTenantNameMapKey)
	tenantNameMap := map[tenants.TenantID]string{}
	for tenantIDString, tenantName := range tenantNameMapRaw {
		tenantNameMap[tenants.TenantID(tenantIDString)] = tenantName
	}
	return tenantNameMap
}
