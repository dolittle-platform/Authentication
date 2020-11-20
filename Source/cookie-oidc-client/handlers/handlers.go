package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"dolittle.io/cookie-oidc-client/configuration"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type Base struct {
	Configuration configuration.Configuration
	OauthConfig   oauth2.Config
	SessionStore  sessions.Store
}

func generateNonce() (string, error) {
	buf := make([]byte, 18)
	_, err := io.ReadAtLeast(rand.Reader, buf, 18)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(buf[0:18]), nil
}

type initiateHandler struct {
	*Base
}

func CreateInitiateHandler(h *Base) (*initiateHandler, error) {
	return &initiateHandler{h}, nil
}

func (h *initiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("return_to")
	if returnTo == "" {
		returnTo = fmt.Sprintf("%s/", h.Configuration.RedirectBaseURL)
	}

	session, err := h.SessionStore.New(r, h.Configuration.SessionStoreName)
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

	http.Redirect(w, r, h.OauthConfig.AuthCodeURL(nonce), http.StatusFound)
}

type callbackHandler struct {
	*Base
}

func CreateCallbackHandler(h *Base) (*callbackHandler, error) {
	return &callbackHandler{h}, nil
}

func (h *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GOT CALLBACK")

	ctx := context.Background()
	oauth2Token, err := h.OauthConfig.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("ACCESS TOKEN", oauth2Token.AccessToken)

	// Do more stuff to contents
	cookie := &http.Cookie{
		Name:    h.Configuration.TokenCookieName,
		Value:   oauth2Token.AccessToken,
		Path:    h.Configuration.TokenCookiePath,
		Expires: time.Now().Add(time.Duration(h.Configuration.TokenCookieExpiresInDays) * 24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, fmt.Sprintf("%s/", h.Configuration.RedirectBaseURL), http.StatusFound)
}
