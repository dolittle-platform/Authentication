package flows

import (
	"github.com/spf13/viper"
	"net/url"
)

const (
	logoutFlowIDQueryParameterKey = "flows.logout.flow_id_query_parameter"
	urlsLoggedOutKey              = "urls.logged_out"

	defaultLogoutFlowIDQueryParameter = "logout_challenge"
)

var (
	defaultLoggedOutRedirect = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "error",
	}
)

type Logout struct{}

func (l *Logout) FlowIDQueryParameter() string {
	if value := viper.GetString(logoutFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultLogoutFlowIDQueryParameter
}

func (l *Logout) LoggedOutRedirect() *url.URL {
	value := viper.GetString(urlsLoggedOutKey)
	if value == "" {
		return defaultLoggedOutRedirect
	}
	url, err := url.Parse(value)
	if err != nil {
		return defaultLoggedOutRedirect
	}
	return url
}
