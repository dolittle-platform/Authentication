package handlers

import (
	"context"
	"log"
	"net/http"
	"time"
)

type CallbackHandler struct {
	*Base
}

func (h *CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received callback")

	ctx := context.Background()
	oauth2Token, err := h.OauthConfig.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Do more stuff to contents
	cookie := &http.Cookie{
		Name:    h.Configuration.Cookie.Name,
		Value:   oauth2Token.AccessToken,
		Path:    h.Configuration.Cookie.Path,
		Expires: time.Now().Add(30 * 24 * time.Hour),
	}
	log.Println("Got access token and setting cookie")
	http.SetCookie(w, cookie)

	log.Println("Redirecting to ", h.Configuration.CallbackRedirectURL)
	http.Redirect(w, r, h.Configuration.CallbackRedirectURL, http.StatusFound)
}
