package viper

import (
	"net/url"

	"dolittle.io/login/server/public"
	"github.com/spf13/viper"
)

const (
	servePortKey                    = "serve.port"
	frontendShowDolittleHeadlineKey = "frontend.showDolittleHeadline"
	frontendApplicationNameKey      = "frontend.applicationName"
	frontendSupportEmailKey         = "frontend.supportEmail"
	urlsErrorKey                    = "urls.error"

	defaultServePort = 8080
)

var (
	defaultErrorRedirect = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "error",
	}
)

type serverConfiguration struct {
	devMode bool
}

func (c *serverConfiguration) Port() int {
	port := viper.GetInt(servePortKey)
	if port == 0 {
		return defaultServePort
	}
	return port
}

func (c *serverConfiguration) DevMode() bool {
	return c.devMode
}

func (c *serverConfiguration) Frontend() public.FrontendConfiguration {
	return public.FrontendConfiguration{
		ShowDolittleHeadline: viper.GetBool(frontendShowDolittleHeadlineKey),
		ApplicationName:      viper.GetString(frontendApplicationNameKey),
		SupportEmail:         viper.GetString(frontendSupportEmailKey),
	}
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
