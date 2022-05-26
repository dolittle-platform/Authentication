package public

import (
	"context"
	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/openid"
	"dolittle.io/pascal/server/handling"
	"go.uber.org/zap"
	"net/http"
)

type LogoutHandler handling.Handler

func NewLogoutHandler(reader cookies.Reader, revoker openid.TokenRevoker, logger *zap.Logger) LogoutHandler {
	return &logout{
		reader:  reader,
		revoker: revoker,
		logger:  logger,
	}
}

type logout struct {
	reader  cookies.Reader
	revoker openid.TokenRevoker
	logger  *zap.Logger
}

func (h *logout) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	token, err := h.reader.ReadTokenCookie(r)
	if err == nil {
		if err := h.revoker.Revoke(token); err != nil {
			return err
		}
	} else {
		h.logger.Debug("no token found in logout request, not revoking token", zap.Error(err))
	}

	w.WriteHeader(200)
	w.Write([]byte("Logging out..."))
	return nil
}
