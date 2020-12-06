package public

import (
	"context"
	"net/http"

	"dolittle.io/login/server/handling"
)

type FrontendHandler handling.Handler

func NewFrontendHandler() FrontendHandler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot/index.html")
	})
	mux.Handle("/.auth/assets/", http.StripPrefix("/.auth/assets/", http.FileServer(http.Dir("wwwroot"))))

	return &frontend{
		mux: mux,
	}
}

type frontend struct {
	mux *http.ServeMux
}

func (f *frontend) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	f.mux.ServeHTTP(w, r)
	return nil
}
