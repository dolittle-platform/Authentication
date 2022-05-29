package viper

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	servePortKey          = "serve.port"
	servePathsInitiateKey = "serve.paths.initiate"
	servePathsCompleteKey = "serve.paths.complete"
	servePathsLogoutKey   = "serve.paths.logout"
	urlsErrorKey          = "urls.error"

	defaultServePort         = 8080
	defaultServeInitiatePath = "/initiate"
	defaultServeCompletePath = "/callback"
	defaultServeLogoutPath   = "/logout"
)

var (
	defaultErrorRedirect = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "error",
	}
)

type serverConfiguration struct{}

func (c *serverConfiguration) Port() int {
	port := viper.GetInt(servePortKey)
	if port == 0 {
		return defaultServePort
	}
	return port
}

func (c *serverConfiguration) InitiatePath() string {
	if path := viper.GetString(servePathsInitiateKey); path != "" {
		return path
	}
	return defaultServeInitiatePath
}

func (c *serverConfiguration) CompletePath() string {
	if path := viper.GetString(servePathsCompleteKey); path != "" {
		return path
	}
	return defaultServeCompletePath
}

func (c *serverConfiguration) LogoutPath() string {
	if path := viper.GetString(servePathsLogoutKey); path != "" {
		return path
	}
	return defaultServeLogoutPath
}

func (c *serverConfiguration) ErrorRedirect() *url.URL {
	value := viper.GetString(urlsErrorKey)
	if value == "" {
		return defaultErrorRedirect
	}
	url, err := url.Parse(value)
	if err != nil {
		return defaultErrorRedirect
	}
	return url
}
