package tenant

import (
	"context"
	"net/http"
	"net/url"

	"dolittle.io/login/flows/tenant"
	"dolittle.io/login/identities/current"
	"dolittle.io/login/server/handling"
)

type InitiateHandler handling.Handler

func NewInitiateHandler(flows tenant.Getter, selecter tenant.Selecter) InitiateHandler {
	return &initiateHandler{
		flows:    flows,
		selecter: selecter,
	}
}

type initiateHandler struct {
	flows    tenant.Getter
	selecter tenant.Selecter
}

func (h *initiateHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.flows.GetTenantFlowFrom(r)
	if err == current.ErrNoUserLoggedIn {
		params := url.Values{}
		params.Add("return_to", r.URL.String())
		http.Redirect(w, r, "/.auth/self-service/login/browser?"+params.Encode(), http.StatusFound)
		return nil
	}
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
