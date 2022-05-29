package logout

import (
	"context"
	"dolittle.io/login/flows/logout"
	"dolittle.io/login/server/handling"
	"net/http"
)

type InitiateHandler handling.Handler

func NewInitiateHandler(getter logout.Getter, accepter logout.Accepter, initiator logout.Initiator) InitiateHandler {
	return &initiateHandler{
		getter:    getter,
		accepter:  accepter,
		initiator: initiator,
	}
}

type initiateHandler struct {
	getter    logout.Getter
	accepter  logout.Accepter
	initiator logout.Initiator
}

func (h *initiateHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	flow, err := h.getter.GetLogoutFlowFrom(r)
	if err == nil {
		redirect, err := h.accepter.AcceptLogoutFlow(ctx, flow)
		if err != nil {
			return err
		}

		http.Redirect(w, r, redirect.String(), http.StatusFound)
		return nil
	} else if err != logout.ErrLogoutChallengeNotFound {
		return err
	}

	redirect, err := h.initiator.InitiateLogout(ctx, r.Cookies())
	if err != nil {
		return err
	}

	http.Redirect(w, r, redirect.String(), http.StatusFound)
	return nil
}
