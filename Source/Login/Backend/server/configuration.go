package server

import (
	"dolittle.io/login/server/handling"
	"dolittle.io/login/server/public"
)

type Configuration interface {
	Port() int

	handling.Configuration
	public.Configuration
}
