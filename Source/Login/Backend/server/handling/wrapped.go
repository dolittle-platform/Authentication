package handling

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type wrappedHandler struct {
	configuration Configuration
	logger        *zap.Logger
	handler       Handler
}

func (h *wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlation := uuid.New().String()
	h.logger.Info("Handling request", zap.String("correlation", correlation), zap.String("path", r.URL.Path))
	ctx := context.TODO()

	defer h.recoverPanic(w, r, correlation)

	err := h.handler.Handle(w, r, ctx)

	if err != nil {
		h.logger.Error("Error during request", zap.String("correlation", correlation), zap.Error(err))
		h.redirectError(w, r, correlation)
		return
	}

	h.logger.Info("Handled request", zap.String("correlation", correlation), zap.String("path", r.URL.Path))
}

func (h *wrappedHandler) recoverPanic(w http.ResponseWriter, r *http.Request, correlation string) {
	if err := recover(); err != nil {
		h.logger.Error("Recovered from request panic", zap.String("correlation", correlation), zap.Reflect("error", err))
		h.redirectError(w, r, correlation)
	}
}

func (h *wrappedHandler) redirectError(w http.ResponseWriter, r *http.Request, correlation string) {
	base := h.configuration.ErrorRedirect().String()
	redirect := fmt.Sprintf("%s?correlation=%s", base, correlation)
	http.Redirect(w, r, redirect, http.StatusFound)
}
