package handlers

import (
	"log"
	"net/http"
	"net/url"

	"dolittle.io/cookie-oidc-client/utils"
)

type initiateHandler struct {
	*Base
	nonceGenerator utils.NonceGenerator
}

type ReturnToGetter interface {
	GetReturnToUrl() url.URL
}

type SessionThingy interface {
	WriteSession(nonce string, returnTo url.URL, w http.ResponseWriter, r *http.Request) error
	ReadSession(, w http.ResponseWriter, r *http.Request) (nonce string, returnTo url.UR)
}

func NewInitiateHandler(base *Base, factory Factory) http.Handler {
	return &initiateHandler{
		Base:           base,
		nonceGenerator: factory.GetNonceGenerator(),
	}
}

func (h *initiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("return_to")
	if returnTo == "" {
		returnTo = h.Configuration.DefaultReturnTo
	}

	session, err := h.SessionStore.New(r, h.Configuration.SessionStoreName)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := h.nonceGenerator.Generate()
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["nonce"] = nonce
	session.Values["return_to"] = returnTo

	err = session.Save(r, w)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, h.OauthConfig.AuthCodeURL(nonce), http.StatusFound)
}
