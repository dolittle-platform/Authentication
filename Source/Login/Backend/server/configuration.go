package server

import "dolittle.io/login/server/handling"

type Configuration interface {
	Port() int

	handling.Configuration
}
