package hydra

import "net/url"

type Configuration interface {
	AdminEndpoint() *url.URL
}
