package server

import (
	"dolittle.io/pascal/server/handling"
)

type Configuration interface {
	Port() int

	InitiatePath() string
	CompletePath() string
	LogoutPath() string

	handling.Configuration
}
