package viper

import (
	"net/url"

	"github.com/ory/viper"
)

const (
	servePortKey          = "serve.port"
	servePathsFrontendKey = "serve.paths.frontend"
	servePathsConsentKey  = "serve.paths.consent"
	urlsErrorKey          = "urls.error"

	defaultServePort         = 8080
	defaultServeFrontendPath = "/"
	defaultServeConsentPath  = "/consent"
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

func (c *serverConfiguration) FrontendPath() string {
	if path := viper.GetString(servePathsFrontendKey); path != "" {
		return path
	}
	return defaultServeFrontendPath
}

func (c *serverConfiguration) ConsentPath() string {
	if path := viper.GetString(servePathsConsentKey); path != "" {
		return path
	}
	return defaultServeConsentPath
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
