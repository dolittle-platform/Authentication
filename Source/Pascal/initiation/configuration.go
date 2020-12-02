package initiation

import "net/url"

type Configuration interface {
	ReturnToParameter() string
	DefaultReturnTo() *url.URL
	AllowedReturnTo() []*url.URL
}
