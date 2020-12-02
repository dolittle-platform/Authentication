package viper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

const (
	cookiesNameKey     = "name"
	cookiesSecureKey   = "secure"
	cookiesSameSiteKey = "samesite"
	cookiesPathKey     = "path"
	cookiesKey         = "cookies"

	defaultCookiesName         = ".oidc-client-cookies"
	defaultCookiesSameSiteMode = http.SameSiteLaxMode
	defaultCookiesPath         = "/"
)

type cookiesConfiguration struct {
	prefix          string
	defaultName     string
	defaultSameSite http.SameSite
	defaultPath     string
}

func (c *cookiesConfiguration) Name() string {
	if name := viper.GetString(fmt.Sprintf("%.%", c.prefix, cookiesNameKey)); name != "" {
		return name
	}
	return c.defaultName
}

func (c *cookiesConfiguration) Secure() bool {
	key := fmt.Sprintf("%.%", c.prefix, cookiesSecureKey)
	if viper.IsSet(key) {
		return viper.GetBool(key)
	}
	return true
}

func (c *cookiesConfiguration) SameSite() http.SameSite {
	mode := viper.GetString(fmt.Sprintf("%.%", c.prefix, cookiesSameSiteKey))
	switch {
	case strings.EqualFold("strict", mode):
		return http.SameSiteStrictMode
	case strings.EqualFold("lax", mode):
		return http.SameSiteLaxMode
	case strings.EqualFold("none", mode):
		return http.SameSiteNoneMode
	default:
		return c.defaultSameSite
	}
}

func (c *cookiesConfiguration) Path() string {
	if path := viper.GetString(fmt.Sprintf("%.%", c.prefix, cookiesPathKey)); path != "" {
		return path
	}
	return c.defaultPath
}
