package config

import (
	"net/http"
	"net/url"
	"strings"
)

func getHttpClientFor(issuer *url.URL, query url.Values) *http.Client {
	return &http.Client{
		Transport: &roundTripperWithWellKnownRequestQuery{
			issuer: issuer,
			query:  query,
		},
	}
}

type roundTripperWithWellKnownRequestQuery struct {
	issuer *url.URL
	query  url.Values
}

func (r *roundTripperWithWellKnownRequestQuery) RoundTrip(request *http.Request) (*http.Response, error) {
	wellKnownEndpoint := strings.TrimSuffix(r.issuer.String(), "/") + "/.well-known/openid-configuration"

	if request.URL.String() == wellKnownEndpoint {
		request.URL.RawQuery = r.query.Encode()
	}

	return http.DefaultTransport.RoundTrip(request)
}
