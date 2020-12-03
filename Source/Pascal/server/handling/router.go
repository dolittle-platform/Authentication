package handling

import (
	"net/http"

	"go.uber.org/zap"
)

type Router interface {
	Handle(pattern string, handler Handler)
	http.Handler
}

func NewRouter(configuration Configuration, logger *zap.Logger) Router {
	return &router{
		configuration: configuration,
		logger:        logger,
		mux:           http.NewServeMux(),
	}
}

type router struct {
	configuration Configuration
	logger        *zap.Logger
	mux           *http.ServeMux
}

func (r *router) Handle(pattern string, handler Handler) {
	r.mux.Handle(pattern, &wrappedHandler{
		configuration: r.configuration,
		logger:        r.logger,
		handler:       handler,
	})
}

func (r *router) ServeHTTP(w http.ResponseWriter, re *http.Request) {
	r.mux.ServeHTTP(w, re)
}
