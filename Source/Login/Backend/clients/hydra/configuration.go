package hydra

import "net/url"

type Configuration interface {
	Endpoint() *url.URL
}
