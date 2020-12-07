package tenant

import (
	"context"
	"net/http"

	"dolittle.io/login/flows/tenant"
	"dolittle.io/login/server/handling"
)

type SelectHandler handling.Handler

func NewSelectHandler(flows tenant.Getter, selecter tenant.Selecter) SelectHandler {
	return &selectHandler{
		flows:    flows,
		selecter: selecter,
	}
}

type selectHandler struct {
	flows    tenant.Getter
	selecter tenant.Selecter
}

func (h *selectHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.flows.GetTenantFlowFrom(r)
	if err != nil {
		return err
	}

	redirect, err := h.selecter.SelectTenantFrom(flow, r)
	if err != nil {
		return err
	}

	http.Redirect(w, r, redirect.String(), http.StatusFound)
	return nil
}
