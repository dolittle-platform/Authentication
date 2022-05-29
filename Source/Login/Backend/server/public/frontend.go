package public

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"net/http/httputil"

	"dolittle.io/login/server/handling"
)

type FrontendHandler handling.Handler

func NewFrontendHandler(configuration Configuration) (FrontendHandler, error) {
	if configuration.DevMode() {
		return &frontend{
			handler: createLocalhostProxyHandler(),
		}, nil
	}

	handler, err := createStaticFileHandler(configuration)

	return &frontend{handler}, err
}

type frontend struct {
	handler http.Handler
}

func (f *frontend) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	f.handler.ServeHTTP(w, r)
	return nil
}

func renderIndexTemplateWithConfiguration(configuration FrontendConfiguration) ([]byte, error) {
	indexTemplate, err := template.ParseFiles("wwwroot/index.html")
	if err != nil {
		return nil, err
	}

	indexRendered := bytes.Buffer{}

	if err := indexTemplate.Execute(&indexRendered, configuration); err != nil {
		return nil, err
	}

	return indexRendered.Bytes(), nil
}

func createStaticFileHandler(configuration Configuration) (http.Handler, error) {
	mux := http.NewServeMux()

	index, err := renderIndexTemplateWithConfiguration(configuration.Frontend())
	if err != nil {
		return nil, err
	}

	mux.HandleFunc("/.auth/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(index)
	})
	mux.Handle("/.auth/assets/", http.StripPrefix("/.auth/assets/", http.FileServer(http.Dir("wwwroot"))))
	return mux, nil
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
