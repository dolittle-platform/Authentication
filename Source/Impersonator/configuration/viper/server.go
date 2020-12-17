package viper

import "github.com/spf13/viper"

const (
	serveProxyPortKey       = "serve.proxy.port"
	serveProxyPathPrefixKey = "serve.proxy.path_prefix"
	serveMetricsPortKey     = "serve.metrics.port"

	defaultServeProxyPort       = 8080
	defaultServeProxyPathPrefix = "/"
	defaultServeMetricsPort     = 9700
)

type serverConfiguration struct{}

func (c *serverConfiguration) ProxyPort() int {
	if port := viper.GetInt(serveProxyPortKey); port != 0 {
		return port
	}
	return defaultServeProxyPort
}

func (c *serverConfiguration) ProxyPathPrefix() string {
	if pathPrefix := viper.GetString(serveProxyPathPrefixKey); pathPrefix != "" {
		return pathPrefix
	}
	return defaultServeProxyPathPrefix
}

func (c *serverConfiguration) MetricsPort() int {
	if port := viper.GetInt(serveMetricsPortKey); port != 0 {
		return port
	}
	return defaultServeMetricsPort
}
