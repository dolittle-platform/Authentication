package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"dolittle.io/cookie-oidc-client/configuration"
	"dolittle.io/cookie-oidc-client/handlers"
	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

func getConfiguration() *configuration.Configuration {
	defaultConfig := configuration.GetDefaults()
	err := configuration.Setup(&defaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	config, err := configuration.Read()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	config := getConfiguration()
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, fmt.Sprintf("%s:%d/", config.ProviderHost, config.ProviderPort))
	if err != nil {
		log.Fatal(err)
	}

	handlerBase := &handlers.Base{
		OauthConfig: oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  fmt.Sprintf("%s/.auth/callback", config.RedirectBaseURL),
			Scopes:       config.Scopes,
		},
		SessionStore: sessions.NewCookieStore([]byte("super-secret-value")),
	}
	initiateHandler, err := handlers.CreateInitiateHandler(handlerBase)
	http.Handle("/initiate/", initiateHandler)
	callbackHandler, err := handlers.CreateCallbackHandler(handlerBase)
	http.Handle("/callback/", callbackHandler)

	http.Handle("/", http.FileServer(http.Dir("spa")))

	log.Println(fmt.Sprintf("Listening on port %d", config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
