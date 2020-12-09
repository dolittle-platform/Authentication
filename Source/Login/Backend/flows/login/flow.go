package login

import (
	"net/url"

	"dolittle.io/login/providers"
)

type FlowID string

type Flow struct {
	ID               FlowID
	Forced           bool
	FormCSRFToken    string
	FormSubmitAction *url.URL
	FormSubmitMethod string

	Providers []providers.Provider
}
