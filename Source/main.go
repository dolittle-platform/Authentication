package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type handler struct {
	oidcConfig  oidc.Config
	oauthConfig oauth2.Config
}

type initiateHandler struct {
	*handler
}

func (h *initiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling something here")
	http.Redirect(w, r, h.oauthConfig.AuthCodeURL("at least eight characters long"), http.StatusFound)
}

type callbackHandler struct {
	*handler
}

func (h *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GOT CALLBACK")

	ctx := context.Background()
	oauth2Token, err := h.oauthConfig.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		// handle error
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		// handle missing token
	}

	// Do more stuff to contents

	w.Header().Add("Cookie-dolittle", rawIDToken)

	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}

func main() {
	fmt.Println("Hello world")

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}

	h := &handler{
		oidcConfig: oidc.Config{
			ClientID: "a5fbb152-4e2a-478a-8744-6f4ca7a65f63",
		},
		oauthConfig: oauth2.Config{
			ClientID:     "a5fbb152-4e2a-478a-8744-6f4ca7a65f63",
			ClientSecret: "FfQvOn5qu0GLeSke24a.c2ZJ4G",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://localhost:8888/callback",
			Scopes:       []string{oidc.ScopeOpenID},
		},
	}

	http.Handle("/initiate", &initiateHandler{h})
	http.Handle("/callback", &callbackHandler{h})

	http.Handle("/", http.FileServer(http.Dir("spa")))

	http.ListenAndServe(":8888", nil)
}
