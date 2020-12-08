package login

import (
	"context"
	"net/http"

	"dolittle.io/login/flows/login"
	"dolittle.io/login/server/handling"
	"dolittle.io/login/server/httputils"
)

type GetHandler handling.Handler

func NewGetHandler(getter login.Getter) GetHandler {
	return &getHandler{
		logins: getter,
	}
}

type getHandler struct {
	logins login.Getter
}

// Handles GET /.auth/self-service/login/flows?id=<id> and provides a list of login providers from Kratos
// for that specific flow
func (h *getHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.logins.GetLoginFlowFrom(r)
	if err != nil {
		return err
	}

	return httputils.WriteJson(w, flow, http.StatusFound)
}
