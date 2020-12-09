package kratos

import "net/url"

type Configuration interface {
	PublicEndpoint() *url.URL
}
