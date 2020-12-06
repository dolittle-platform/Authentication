package consent

import (
	"context"
	"net/http"

	"dolittle.io/login/flows/consent"
	"dolittle.io/login/server/handling"
)

type InitiateHandler handling.Handler

func NewInitiateHandler() InitiateHandler {
	return &initiateHandler{}
}

type initiateHandler struct {
	getter   consent.Getter
	accepter consent.Accepter
}

func (h *initiateHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.getter.GetConsentFlowFrom(r)
	if err != nil {
		return err
	}

	redirect, err := h.accepter.AcceptConsentFlow(flow)
	if err != nil {
		return err
	}

	http.Redirect(w, r, redirect.String(), http.StatusFound)
	return nil
}
