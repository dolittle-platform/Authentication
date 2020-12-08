package tenant

import (
	"context"
	"net/http"

	"dolittle.io/login/flows/tenant"
	"dolittle.io/login/server/handling"
	"dolittle.io/login/server/httputils"
)

type GetHandler handling.Handler

func NewGetHandler(flows tenant.Getter) GetHandler {
	return &getHandler{
		flows: flows,
	}
}

type getHandler struct {
	flows tenant.Getter
}

// /.auth/self-service/tenant/flows
func (h *getHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.flows.GetTenantFlowFrom(r)
	if err != nil {
		return err
	}

	return httputils.WriteJson(w, flow, http.StatusOK)
}
