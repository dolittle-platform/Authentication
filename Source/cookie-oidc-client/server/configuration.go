package server

import "dolittle.io/cookie-oidc-client/server/handling"

type Configuration interface {
	Port() int

	InitiatePath() string
	CompletePath() string

	handling.Configuration
}
