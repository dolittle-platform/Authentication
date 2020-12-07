package public

import (
	"context"
	"net/http"
	"net/http/httputil"

	"dolittle.io/login/server/handling"
)

type FrontendHandler handling.Handler

func NewFrontendHandler(configuration Configuration) FrontendHandler {
	if configuration.DevMode() {
		return &frontend{
			handler: createLocalhostProxyHandler(),
		}
	}
	return &frontend{
		handler: createStaticFileHandler(),
	}
}

type frontend struct {
	handler http.Handler
}

func (f *frontend) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	f.handler.ServeHTTP(w, r)
	return nil
}

func createStaticFileHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/.auth/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot/index.html")
	})
	mux.Handle("/.auth/assets/", http.StripPrefix("/.auth/assets/", http.FileServer(http.Dir("wwwroot"))))
	return mux
}

func createLocalhostProxyHandler() http.Handler {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = "localhost:8091"
			r.Host = "localhost:8091"
		},
	}
}
