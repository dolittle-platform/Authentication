package viper

import (
	"dolittle.io/pascal/redirects"
	"net/url"

	"github.com/spf13/viper"
)

const (
	urlsReturnQueryParameterKey = "urls.return.query_parameter"
	urlsLoginReturnDefaultKey   = "urls.return.default.login"
	urlsLogoutReturnDefaultKey  = "urls.return.default.logout"
	urlsReturnAllowedKey        = "urls.return.allowed"
	urlsReturnModeKey           = "urls.return.mode"

	defaultReturnToParameter = "return_to"
	defaultReturnMode        = redirects.MatchModeStrict
)

var (
	defaultLoginReturnTo = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "return",
	}
	defaultLogoutReturnTo = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "return",
	}
)

type redirectsConfiguration struct{}

func (c *redirectsConfiguration) ReturnToParameter() string {
	if parameter := viper.GetString(urlsReturnQueryParameterKey); parameter != "" {
		return parameter
	}
	return defaultReturnToParameter
}

func (c *redirectsConfiguration) DefaultLoginReturnTo() *url.URL {
	value := viper.GetString(urlsLoginReturnDefaultKey)
	if value == "" {
		return defaultLoginReturnTo
	}
	url, err := url.Parse(value)
	if err != nil {
		return defaultLoginReturnTo
	}
	return url
}

func (c *redirectsConfiguration) DefaultLogoutReturnTo() *url.URL {
	value := viper.GetString(urlsLogoutReturnDefaultKey)
	if value == "" {
		return defaultLogoutReturnTo
	}
	url, err := url.Parse(value)
	if err != nil {
		return defaultLogoutReturnTo
	}
	return url
}

func (c *redirectsConfiguration) AllowedReturnTo() []*url.URL {
	allowed := []*url.URL{
		c.DefaultLoginReturnTo(),
		c.DefaultLogoutReturnTo(),
	}

	for _, value := range viper.GetStringSlice(urlsReturnAllowedKey) {
		if url, err := url.Parse(value); err != nil {
			allowed = append(allowed, url)
		}
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
