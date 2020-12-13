package client

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client http.RoundTripper

func NewClient(configuration Configuration) (Client, error) {
	transport, err := getTransportWithRootCAs(configuration)
	if err != nil {
		return nil, err
	}
	token, err := getServiceAccountToken(configuration)
	if err != nil {
		return nil, err
	}

	return &client{
		transport: transport,
		token:     token,
	}, nil
}

type client struct {
	transport *http.Transport
	token     string
}

func (c *client) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+c.token)
	return c.transport.RoundTrip(r)
}

func getTransportWithRootCAs(configuration Configuration) (*http.Transport, error) {
	data, err := ioutil.ReadFile(configuration.CertificateAuthorityPath())
	if err != nil {
		return nil, err
	}

	certificates := x509.NewCertPool()
	if !certificates.AppendCertsFromPEM(data) {
		return nil, errors.New("invalid ca certs data")
	}

	return &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: certificates,
		},
		ForceAttemptHTTP2: true,
	}, nil
}

func getServiceAccountToken(configuration Configuration) (string, error) {
	data, err := ioutil.ReadFile(configuration.ServiceAccountTokenPath())
	if err != nil {
		return "", err
	}
	return string(data), nil
}
