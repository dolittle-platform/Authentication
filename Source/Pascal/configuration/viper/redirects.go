package viper

import (
	"dolittle.io/pascal/redirects"
	"github.com/spf13/viper"
)

const (
	urlsReturnQueryParameterKey = "urls.return.query_parameter"
	urlsLoginReturnDefaultKey   = "urls.return.default.login"
	urlsLogoutReturnDefaultKey  = "urls.return.default.logout"
	urlsReturnAllowedKey        = "urls.return.allowed"
	urlsReturnModeKey           = "urls.return.mode"

	defaultReturnToParameter = "return_to"
	defaultLoginReturnTo     = "return"
	defaultLogoutReturnTo    = "return"
	defaultReturnMode        = redirects.MatchModeStrict
)

type redirectsConfiguration struct{}

func (c *redirectsConfiguration) ReturnToParameter() string {
	if parameter := viper.GetString(urlsReturnQueryParameterKey); parameter != "" {
		return parameter
	}
	return defaultReturnToParameter
}

func (c *redirectsConfiguration) DefaultLoginReturnTo() string {
	if value := viper.GetString(urlsLoginReturnDefaultKey); value != "" {
		return value
	}
	return defaultLoginReturnTo
}

func (c *redirectsConfiguration) DefaultLogoutReturnTo() string {
	if value := viper.GetString(urlsLogoutReturnDefaultKey); value != "" {
		return value
	}
	return defaultLogoutReturnTo
}

func (c *redirectsConfiguration) AllowedReturnTo() []string {
	allowed := []string{
		c.DefaultLoginReturnTo(),
		c.DefaultLogoutReturnTo(),
	}

	for _, value := range viper.GetStringSlice(urlsReturnAllowedKey) {
		allowed = append(allowed, value)
	}

	return allowed
}

func (c *redirectsConfiguration) ReturnToMatchMode() redirects.MatchMode {
	switch viper.GetString(urlsReturnModeKey) {
	case "prefix":
		return redirects.MatchModePrefix
	default:
		return defaultReturnMode
	}
}
