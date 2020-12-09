package viper

import "github.com/spf13/viper"

const (
	identitiesCurrentUserCookieNameKey = "identities.cookie_name"

	defaultIdentitiesCurrentUserCookieNameValue = "ory_kratos_session"
)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) Cookie() string {
	if value := viper.GetString(identitiesCurrentUserCookieNameKey); value != "" {
		return value
	}
	return defaultIdentitiesCurrentUserCookieNameValue
}
