package redirects

import (
	"net/http"
	"net/url"
)

func GetHostFor(r *http.Request) string {
	request := url.URL{
		Scheme: "http",
		Host:   r.Host,
	}

	if r.TLS != nil {
		request.Scheme = "https"
	}

	for _, scheme := range r.Header.Values("x-forwarded-proto") {
		request.Scheme = scheme
	}

	for _, host := range r.Header.Values("x-forwarded-host") {
		request.Host = host
	}

	return request.String()
}

func GetAbsoluteUrlFor(r *http.Request, suffix string) (*url.URL, error) {
	return url.Parse(GetHostFor(r) + suffix)
}
