package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

type InitiateHandler struct {
	*Base
}

func (h *InitiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("return_to")
	if returnTo == "" {
		returnTo = h.Configuration.DefaultReturnTo
	}

	session, err := h.SessionStore.New(r, h.Configuration.SessionStoreName)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := generateNonce()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["nonce"] = nonce
	session.Values["return_to"] = returnTo

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, h.OauthConfig.AuthCodeURL(nonce), http.StatusFound)
}

type CallbackHandler struct {
	*Base
}

func (h *CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("GOT CALLBACK")

	ctx := context.Background()
	oauth2Token, err := h.OauthConfig.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("ACCESS TOKEN", oauth2Token.AccessToken)

	// Do more stuff to contents
	cookie := &http.Cookie{
		Name:    h.Configuration.Cookie.Name,
		Value:   oauth2Token.AccessToken,
		Path:    h.Configuration.Cookie.Path,
		Expires: time.Now().Add(30 * 24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, h.Configuration.CallbackRedirectURL, http.StatusFound)
}
