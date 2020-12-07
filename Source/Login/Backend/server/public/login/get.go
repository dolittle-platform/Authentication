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

func (h *getHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.logins.GetLoginFlowFrom(r)
	if err != nil {
		return err
	}

	return httputils.WriteJson(w, flow, http.StatusFound)
}
