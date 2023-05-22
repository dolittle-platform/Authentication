package tenant

import (
	"context"
	"net/http"
	"net/url"

	"dolittle.io/login/flows/tenant"
	"dolittle.io/login/identities/current"
	"dolittle.io/login/server/handling"
)

// Handles initiating the tenant selection
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

// Basically this handle just makes sure we are logged in and allows us to speedline the process for single-tenant users
// Handle handles GET /.auth/self-service/tenant/browser which is the Hydra login endpoint.
// Checks if have a login flow from Hydra first with GetTenantFlowFrom, which also checks if we are logged in Kraots
// If we aren't logged in yet, we start the Kratos login flow.
// If we are logged in, add the login challenge to the url and redirect us to select a tenant with the same
// Hydra login_challenge query in the url again
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

	if len(flow.User.Tenants) == 0 {
		http.Redirect(w, r, "/.auth/no-tenant", http.StatusFound)
		return nil
	}

	if len(flow.User.Tenants) == 1 {
		redirect, err := h.selecter.SelectTenant(ctx, flow, flow.User.Tenants[0].ID)
		if err != nil {
			return err
		}

		http.Redirect(w, r, redirect.String(), http.StatusFound)
		return nil
	}

	params := url.Values{}
	params.Add("login_challenge", string(flow.ID))
	http.Redirect(w, r, "/.auth/select-tenant?"+params.Encode(), http.StatusFound)
	return nil
}
