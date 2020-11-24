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
	provider, err := oidc.NewProvider(ctx, fmt.Sprintf("%s:%d/", config.Provider.Host, config.Provider.Port))
	if err != nil {
		log.Fatal(err)
	}

	handlerBase := &handlers.Base{
		OauthConfig: oauth2.Config{
			ClientID:     config.Provider.ClientID,
			ClientSecret: config.Provider.ClientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  config.Provider.RedirectURL,
			Scopes:       config.Provider.Scopes,
		},
		SessionStore:  sessions.NewCookieStore([]byte("super-secret-value")),
		Configuration: *config,
	}
	http.Handle("/initiate/", &handlers.InitiateHandler{Base: handlerBase})
	http.Handle("/callback/", &handlers.CallbackHandler{Base: handlerBase})

	log.Println(fmt.Sprintf("Listening on port %d", config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
