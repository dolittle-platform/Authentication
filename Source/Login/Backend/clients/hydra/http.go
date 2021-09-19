package hydra

import "net/http"

type clientWithDefaultHeaders struct {
	*http.Client
	Header http.Header
}

func newClientWithDefaultHeaders() *clientWithDefaultHeaders {
	header := make(http.Header)
	return &clientWithDefaultHeaders{
		Client: &http.Client{
			Transport: &transportWithDefaultHeaders{
				Header:    header,
				Transport: http.DefaultTransport,
			},
		},
		Header: header,
	}
}

type transportWithDefaultHeaders struct {
	Header    http.Header
	Transport http.RoundTripper
}

func (t *transportWithDefaultHeaders) RoundTrip(r *http.Request) (*http.Response, error) {
	for name, value := range t.Header {
		r.Header[name] = append(r.Header[name], value...)
	}
	return t.Transport.RoundTrip(r)
}
