package handlers

import (
	"dolittle.io/cookie-oidc-client/configuration"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type Base struct {
	Configuration configuration.Configuration
	OauthConfig   oauth2.Config
	SessionStore  sessions.Store
}
