package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxyTo(proxy string) http.Handler {
	address, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(address)
}

func strip(path string, handler http.Handler) http.Handler {
	return http.StripPrefix(path, handler)
}

func main() {
	// Studio (Oathkeeper)
	http.Handle("/", proxyTo("http://localhost:8001/"))
	// Hydra
	http.Handle("/.auth/.well-known/", strip("/.auth/", proxyTo("http://localhost:4444/")))
	http.Handle("/.auth/oauth2/", strip("/.auth/", proxyTo("http://localhost:4444/")))
	// Kratos
	http.Handle("/.auth/self-service/", strip("/.auth/", proxyTo("http://localhost:4433/")))

	// Pascal
	http.Handle("/.auth/cookies/", proxyTo("http://localhost:8002/"))

	// Login
	http.Handle("/.auth/", proxyTo("http://localhost:8090/"))
	http.Handle("/.auth/self-service/tenant/", proxyTo("http://localhost:8090/"))
	http.Handle("/.auth/self-service/login/flows", proxyTo("http://localhost:8090/"))

	fmt.Println("Starting reverse proxy on :8080")
	http.ListenAndServe(":8080", nil)
}
