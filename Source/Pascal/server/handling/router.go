package handling

import (
	"go.uber.org/zap"
	"net/http"
)

type Router interface {
	Handle(pattern string, handler Handler)
	http.Handler
}

func NewRouter(configuration Configuration, logger *zap.Logger) Router {
	return &router{
		allowedHosts:  configuration.AllowedHosts(),
		errorRedirect: configuration.ErrorRedirect(),
		logger:        logger,
		mux:           http.NewServeMux(),
	}
}

type router struct {
	allowedHosts  []string
	errorRedirect string
	logger        *zap.Logger
	mux           *http.ServeMux
}

func (r *router) Handle(pattern string, handler Handler) {
	r.mux.Handle(pattern, &wrappedHandler{
		errorRedirect: r.errorRedirect,
		logger:        r.logger,
		handler:       handler,
	})
}

func (r *router) ServeHTTP(w http.ResponseWriter, re *http.Request) {
	if !r.hostIsAllowed(re) {
		r.logger.Warn("the requested host is not allowed", zap.String("host", re.Host))
		http.NotFound(w, re)
		return
	}

	r.mux.ServeHTTP(w, re)
}

func (r *router) hostIsAllowed(re *http.Request) bool {
	for _, allowed := range r.allowedHosts {
		if re.Host == allowed {
			return true
		}
	}

	return false
}
