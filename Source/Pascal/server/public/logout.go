package public

import (
	"context"
	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/logout"
	"dolittle.io/pascal/server/handling"
	"go.uber.org/zap"
	"net/http"
)

type LogoutHandler handling.Handler

func NewLogoutHandler(crumbler cookies.Crumbler, parser logout.Parser, initiator logout.Initiator, logger *zap.Logger) LogoutHandler {
	return &logoutHandler{
		crumbler:  crumbler,
		parser:    parser,
		initiator: initiator,
		logger:    logger,
	}
}

type logoutHandler struct {
	crumbler  cookies.Crumbler
	parser    logout.Parser
	initiator logout.Initiator
	logger    *zap.Logger
}

func (h *logoutHandler) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	err := h.crumbler.DestroyTokenCookie(w)
	if err != nil {
		return err
	}

	request, err := h.parser.ParseFrom(r)
	if err != nil {
		return err
	}

	redirect, err := h.initiator.Initiate(request)
	if err != nil {
		return err
	}

	http.Redirect(w, r, redirect.String(), http.StatusFound)
	return nil
}
