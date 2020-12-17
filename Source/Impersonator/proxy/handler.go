package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"dolittle.io/impersonator/proxy/client"
	"dolittle.io/impersonator/proxy/context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler http.Handler

func NewHandler(creator context.Creator, modifier Modifier, client client.Client, logger *zap.Logger) Handler {
	h := &handler{
		creator:  creator,
		modifier: modifier,
		logger:   logger,
		contexts: make(map[*http.Request]*context.Context),
	}
	h.proxy = &httputil.ReverseProxy{
		Transport:      client,
		Director:       h.director,
		ModifyResponse: h.modifyResponse,
		ErrorHandler:   h.errorHandler,
		ErrorLog:       newZapLogWriter(logger),
	}
	return h
}

type handler struct {
	proxy    *httputil.ReverseProxy
	creator  context.Creator
	modifier Modifier
	logger   *zap.Logger

	contexts map[*http.Request]*context.Context
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlation := uuid.New().String()
	h.logger.Debug("Handling request", zap.String("correlation", correlation), zap.String("path", r.URL.Path))

	defer h.recoverPanic(w, r, correlation)

	r.Header.Add("correlation", correlation)
	h.proxy.ServeHTTP(w, r)
}

func (h *handler) recoverPanic(w http.ResponseWriter, r *http.Request, correlation string) {
	if panic := recover(); panic != nil {
		if err, ok := panic.(error); ok {
			h.logger.Error("Recovered from request panic", zap.String("correlation", correlation), zap.Error(err))
		} else {
			h.logger.Error("Recovered from request panic", zap.String("correlation", correlation), zap.Any("err", panic))
		}
		w.WriteHeader(http.StatusInternalServerError)
		totalRequestsServed.WithLabelValues(r.Method, "500").Inc()
	}
}

func (h *handler) director(req *http.Request) {
	ctx, err := h.creator.CreateFor(req)
	if err != nil {
		panic(err)
	}
	if err = h.modifier.ModifyRequest(ctx, req); err != nil {
		panic(err)
	}
	h.contexts[req] = ctx
}

func (h *handler) modifyResponse(res *http.Response) error {
	if res.Request == nil {
		return ErrNoRequestSetInResponse
	}
	ctx, ok := h.contexts[res.Request]
	if !ok {
		return ErrNoContextStoredForRequest
	}
	delete(h.contexts, res.Request)
	err := h.modifier.ModifyResponse(ctx, res.Request, res)
	totalRequestsServed.WithLabelValues(res.Request.Method, fmt.Sprint(res.StatusCode)).Inc()
	return err
}

func (h *handler) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if correlation := r.Header.Get("correlation"); correlation != "" {
		h.logger.Error("Error during request handling", zap.String("correlation", correlation), zap.Error(err))
	} else {
		h.logger.Error("Error during request handling", zap.Error(err))
	}
	w.WriteHeader(http.StatusInternalServerError)
}
