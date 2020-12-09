package handling

import "net/url"

type Configuration interface {
	ErrorRedirect() *url.URL
}
