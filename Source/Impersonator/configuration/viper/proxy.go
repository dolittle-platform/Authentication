package viper

import (
	"net/url"

	"github.com/spf13/viper"
)

const (
	proxyApiServerURLKey                      = "proxy.url"
	proxyApiServerTokenPathKey                = "proxy.token_path"
	proxyApiServerCertificateAuthorityPathKey = "proxy.ca_path"

	defaultProxyApiServerTokenPath                = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	defaultProxyApiServerCertificateAuthorityPath = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

var (
	defaultProxyApiServerURL = &url.URL{
		Scheme: "https",
		Host:   "kubernetes.default.svc.cluster.local",
	}
)

type proxyConfiguration struct{}

func (c *proxyConfiguration) APIServerURL() *url.URL {
	if value := viper.GetString(proxyApiServerURLKey); value != "" {
		if address, err := url.Parse(value); err == nil {
			return address
		}
	}
	return defaultProxyApiServerURL
}

func (c *proxyConfiguration) ServiceAccountTokenPath() string {
	if path := viper.GetString(proxyApiServerTokenPathKey); path != "" {
		return path
	}
	return defaultProxyApiServerTokenPath
}

func (c *proxyConfiguration) CertificateAuthorityPath() string {
	if path := viper.GetString(proxyApiServerCertificateAuthorityPathKey); path != "" {
		return path
	}
	return defaultProxyApiServerCertificateAuthorityPath
}
