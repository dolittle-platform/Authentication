package public

import (
	"context"
	"net/http"

	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/server/handling"
	"dolittle.io/cookie-oidc-client/sessions"
)

type InitiateHandler handling.Handler

func NewInitiateHandler(parser initiation.Parser, initiator initiation.Initiatior, writer sessions.Writer) InitiateHandler {
	return &initiate{
		parser:    parser,
		initiator: initiator,
		writer:    writer,
	}
}

type initiate struct {
	parser    initiation.Parser
	initiator initiation.Initiatior
	writer    sessions.Writer
}

func (h *initiate) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	request, err := h.parser.ParseFrom(r)
	if err != nil {
		return err
	}

	session, redirect, err := h.initiator.Initiate(request)
	if err != nil {
		return err
	}

	if err := h.writer.WriteTo(session, r, w); err != nil {
		return err
	}

	http.Redirect(w, r, string(redirect), http.StatusFound)
	return nil
}
