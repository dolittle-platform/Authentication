package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"

	hydra "github.com/ory/hydra-client-go/client"
	admin "github.com/ory/hydra-client-go/client/admin"
	models "github.com/ory/hydra-client-go/models"

	kratos "github.com/ory/kratos-client-go/client"
	public "github.com/ory/kratos-client-go/client/public"

	runtimeClient "github.com/go-openapi/runtime/client"
)

type handler struct {
	oidcConfig   oidc.Config
	oauthConfig  oauth2.Config
	sessionStore sessions.Store
	hydraClient  *hydra.OryHydra
	kratosClient *kratos.OryKratos
}

type initiateHandler struct {
	*handler
}

func generateNonce() (string, error) {
	buf := make([]byte, 18)
	_, err := io.ReadAtLeast(rand.Reader, buf, 18)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(buf[0:18]), nil
}

func (h *initiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("return_to")
	if returnTo == "" {
		returnTo = "http://localhost:8080/"
	}

	session, err := h.sessionStore.New(r, "dolittle-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := generateNonce()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["nonce"] = nonce
	session.Values["return_to"] = returnTo

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, h.oauthConfig.AuthCodeURL(nonce), http.StatusFound)
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

type selectedHandler struct {
	*handler
}

func (h *selectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookieVal := cookie.String()

	whoami, err := h.kratosClient.Public.Whoami(public.NewWhoamiParams().WithCookie(&cookieVal), runtimeClient.PassThroughAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	identity := whoami.Payload.Identity
	subject := string(identity.ID)
	traits := identity.Traits.(map[string]interface{})
	log.Println("IDENTITY", identity, subject, traits)

	// Verify that selected tenant is actually in list of tenants
	body := &models.AcceptLoginRequest{
		Subject: &subject,
		Context: struct {
			email  string
			tenant string
		}{
			email:  traits["email"].(string),
			tenant: r.PostFormValue("tenant"),
		},
		Remember: false,
	}
	params := admin.NewAcceptLoginRequestParams().WithLoginChallenge(r.PostFormValue("login_challenge")).WithBody(body)
	acceptLogin, err := h.hydraClient.Admin.AcceptLoginRequest(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, *acceptLogin.Payload.RedirectTo, http.StatusFound)
}

type consentHandler struct {
	*handler
}

func (h *consentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}

	hydraAdminURL, _ := url.Parse("http://localhost:4445")
	hydraAdmin := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{
		Schemes:  []string{hydraAdminURL.Scheme},
		Host:     hydraAdminURL.Host,
		BasePath: hydraAdminURL.Path,
	})

	kratosPublicURL, _ := url.Parse("http://localhost:8080/.ory/kratos/public/")
	kratosPublic := kratos.NewHTTPClientWithConfig(nil, &kratos.TransportConfig{
		Schemes:  []string{kratosPublicURL.Scheme},
		Host:     kratosPublicURL.Host,
		BasePath: kratosPublicURL.Path,
	})

	h := &handler{
		oidcConfig: oidc.Config{
			ClientID: "do",
		},
		oauthConfig: oauth2.Config{
			ClientID:     "do",
			ClientSecret: "little",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://localhost:8080/.auth/callback/",
			Scopes:       []string{oidc.ScopeOpenID},
		},
		sessionStore: sessions.NewCookieStore([]byte("super-secret-value")),
		hydraClient:  hydraAdmin,
		kratosClient: kratosPublic,
	}

	http.Handle("/initiate/", &initiateHandler{h})
	http.Handle("/selected-tenant/", &selectedHandler{h})
	http.Handle("/consent/", &consentHandler{h})
	http.Handle("/callback", &callbackHandler{h})

	http.Handle("/", http.FileServer(http.Dir("spa")))

	log.Println("Listening on http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
