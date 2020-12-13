package proxy

import (
	"net/url"

	"dolittle.io/impersonator/proxy/client"
)

type Configuration interface {
	APIServerURL() *url.URL
	client.Configuration
}
