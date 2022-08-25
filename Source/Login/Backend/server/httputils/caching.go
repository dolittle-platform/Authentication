package httputils

import "net/http"

func WithCacheControl(directive string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Cache-Control", directive)
		next.ServeHTTP(rw, r)
	})
}
