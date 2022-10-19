package viper

import (
	"github.com/spf13/viper"
)

const (
	servePortKey          = "serve.port"
	serveHostsKey         = "serve.hosts"
	servePathsInitiateKey = "serve.paths.initiate"
	servePathsCompleteKey = "serve.paths.complete"
	servePathsLogoutKey   = "serve.paths.logout"
	urlsErrorKey          = "urls.error"

	defaultServePort         = 8080
	defaultServeInitiatePath = "/initiate"
	defaultServeCompletePath = "/callback"
	defaultServeLogoutPath   = "/logout"
	defaultErrorRedirect     = "/error"
)

var (
	defaultHosts = []string{"localhost:8080"}
)

type serverConfiguration struct{}

func (c *serverConfiguration) Port() int {
	port := viper.GetInt(servePortKey)
	if port == 0 {
		return defaultServePort
	}
	return port
}

func (c *serverConfiguration) AllowedHosts() []string {
	if hosts := viper.GetStringSlice(serveHostsKey); len(hosts) > 0 {
		return hosts
	}

	return defaultHosts
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

func (c *serverConfiguration) ErrorRedirect() string {
	if value := viper.GetString(urlsErrorKey); value != "" {
		return value
	}
	return defaultErrorRedirect
}
