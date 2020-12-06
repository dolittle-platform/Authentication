package tenant

import (
	"context"
	"net/http"

	"dolittle.io/login/flows/tenant"
	"dolittle.io/login/server/handling"
)

type InitiateHandler handling.Handler

func NewInitiateHandler() InitiateHandler {
	return &initiateHandler{}
}

type initiateHandler struct {
	flows    tenant.Getter
	selecter tenant.Selecter
}

func (h *initiateHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.flows.GetTenantFlowFrom(r)
	if err != nil {
		return err
	}

	if len(flow.AvailableTenants) == 1 {
		redirect, err := h.selecter.SelectTenant(flow, flow.AvailableTenants[0])
		if err != nil {
			return err
		}

		http.Redirect(w, r, redirect.String(), http.StatusFound)
		return nil
	}

	http.Redirect(w, r, "/.auth/select-tenant", http.StatusFound)
	return nil
}
