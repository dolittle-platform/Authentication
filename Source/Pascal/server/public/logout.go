package public

import (
	"context"
	"dolittle.io/pascal/server/handling"
	"net/http"
)

type LogoutHandler handling.Handler

func NewLogoutHandler() LogoutHandler {
	return &logout{}
}

type logout struct {
}

func (h *logout) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	w.WriteHeader(200)
	w.Write([]byte("Logging out..."))
	return nil
}
