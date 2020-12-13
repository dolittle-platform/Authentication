package server

type Configuration interface {
	ProxyPort() int
	ProxyPathPrefix() string
	MetricsPort() int
}
