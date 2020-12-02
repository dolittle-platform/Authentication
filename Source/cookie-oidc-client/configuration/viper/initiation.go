package viper

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	urlsReturnQueryParameterKey = "urls.return.query_parameter"
	urlsReturnDefaultKey        = "urls.return.default"
	urlsReturnAllowedKey        = "urls.return.allowed"

	defaultReturnToParameter = "return_to"
)

var (
	defaultReturnTo = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "return",
	}
)

type initiationConfiguration struct{}

func (c *initiationConfiguration) ReturnToParameter() string {
	if parameter := viper.GetString(urlsReturnQueryParameterKey); parameter != "" {
		return parameter
	}
	return defaultReturnToParameter
}

func (c *initiationConfiguration) DefaultReturnTo() *url.URL {
	value := viper.GetString(urlsReturnDefaultKey)
	if value == "" {
		return defaultReturnTo
	}
	url, err := url.Parse(value)
	if err != nil {
		return defaultReturnTo
	}
	return url
}

func (c *initiationConfiguration) AllowedReturnTo() []*url.URL {
	allowed := []*url.URL{c.DefaultReturnTo()}

	for _, value := range viper.GetStringSlice(urlsReturnAllowedKey) {
		if url, err := url.Parse(value); err != nil {
			allowed = append(allowed, url)
		}
	}

	return allowed
}
