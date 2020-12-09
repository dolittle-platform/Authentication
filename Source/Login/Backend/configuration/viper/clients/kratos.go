package clients

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	kratosPublicEndpointKey = "clients.kratos.endpoints.public"
)

var (
	defaultKratosPublicEndpoint = &url.URL{
		Scheme: "http",
		Host:   "localhost:4433",
		Path:   "/",
	}
)

type Kratos struct{}

func (h *Kratos) PublicEndpoint() *url.URL {
	value := viper.GetString(kratosPublicEndpointKey)
	if value == "" {
		return defaultKratosPublicEndpoint
	}
	endpoint, err := url.Parse(value)
	if err != nil {
		return defaultKratosPublicEndpoint
	}
	return endpoint
}
