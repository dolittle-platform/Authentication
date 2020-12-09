package clients

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	hydraAdminEndpointKey = "clients.hydra.endpoints.admin"
)

var (
	defaultHydraAdminEndpoint = &url.URL{
		Scheme: "http",
		Host:   "localhost:4455",
		Path:   "/",
	}
)

type Hydra struct{}

func (h *Hydra) AdminEndpoint() *url.URL {
	value := viper.GetString(hydraAdminEndpointKey)
	if value == "" {
		return defaultHydraAdminEndpoint
	}
	endpoint, err := url.Parse(value)
	if err != nil {
		return defaultHydraAdminEndpoint
	}
	return endpoint
}
