package login

import (
	"dolittle.io/login/flows/forms"
	"dolittle.io/login/providers"
)

type FlowID string

type Flow struct {
	ID        FlowID               `json:"id"`
	Refresh   bool                 `json:"refresh"`
	Form      forms.Form           `json:"form"`
	Providers []providers.Provider `json:"providers"`
}
