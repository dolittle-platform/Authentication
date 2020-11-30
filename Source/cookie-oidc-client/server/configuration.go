package server

import "dolittle.io/cookie-oidc-client/server/handling"

type Configuration interface {
	Port() int

	handling.Configuration
}
