package main

import (
	"context"
	"log"

	"dolittle.io/cookie-oidc-client/configuration"
	handlers "dolittle.io/cookie-oidc-client/handlers"
	"dolittle.io/cookie-oidc-client/urls"
	handlersUrls "dolittle.io/cookie-oidc-client/urls/handlers"
	"dolittle.io/cookie-oidc-client/utils"
	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type Container struct {
	config      *configuration.Configuration
	oauthConfig oauth2.Config

	Server *Server

	TokenGetter            handlers.TokenGetter
	CookieFactory          handlers.CookieFactory
	NonceGenerator         utils.NonceGenerator
	SessionThingy          handlers.SessionThingy
	CallbackRedirectGetter handlersUrls.CallbackRedirectGetter
	ConsentPageGetter      handlersUrls.ConsentPageGetter
	ReturnToGetter         handlersUrls.ReturnToGetter
}

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

func createOauthConfig(config *configuration.Configuration, issuerGetter urls.OIDCProviderIssuerGetter) (oauth2.Config, error) {
	ctx := context.Background()
	issuerUrl, err := issuerGetter.GetOIDCProviderIssuerURL()
	if err != nil {
		return oauth2.Config{}, err
	}

	provider, err := oidc.NewProvider(ctx, issuerUrl.String())
	if err != nil {
		return oauth2.Config{}, err
	}

	return oauth2.Config{
		ClientID:     config.Provider.ClientID,
		ClientSecret: config.Provider.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  config.Provider.RedirectURL,
		Scopes:       config.Provider.Scopes,
	}, nil
}

func (self *Container) Setup() {
	self.config = getConfiguration()
	oauthConfig, err := createOauthConfig(self.config, urls.NewOIDCProviderIssuerGetter(self.config))
	if utils.TryLogIfErrorFatal(err) {
		return
	}
	self.oauthConfig = oauthConfig

	self.Server = NewServer(self.config)

	self.CallbackRedirectGetter = handlersUrls.NewCallbackRedirectGetter(self.config)
	self.ConsentPageGetter = handlersUrls.NewConsentPageGetter(self.oauthConfig)
	self.CookieFactory = handlers.NewOauth2CookieFactory(self.config)
	self.NonceGenerator = utils.NewNonceGenerator()
	returnToGetter, err := handlersUrls.NewReturnToGetter(self.config)
	if utils.TryLogIfErrorFatal(err) {
		return
	}
	self.ReturnToGetter = returnToGetter
	self.SessionThingy = handlers.NewSessionThingy(
		handlers.NewSessions(
			sessions.NewCookieStore([]byte("super-secret-value")), self.config))
	self.TokenGetter = handlers.NewTokenGetter(self.oauthConfig)

}
