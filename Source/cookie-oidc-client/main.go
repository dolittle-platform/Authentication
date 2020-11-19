package main

import (
	"context"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"

	handlers "dolittle.io/cookie-oidc-client/handlers"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}

	handlerBase := &handlers.Handler{
		OauthConfig: oauth2.Config{
			ClientID:     "do",
			ClientSecret: "little",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://localhost:8080/.auth/callback/",
			Scopes:       []string{oidc.ScopeOpenID},
		},
		SessionStore: sessions.NewCookieStore([]byte("super-secret-value")),
	}
	initiateHandler, err := handlers.CreateInitiateHandler(handlerBase)
	http.Handle("/initiate/", initiateHandler)
	callbackHandler, err := handlers.CreateCallbackHandler(handlerBase)
	http.Handle("/callback/", callbackHandler)

	http.Handle("/", http.FileServer(http.Dir("spa")))

	log.Println("Listening on http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
