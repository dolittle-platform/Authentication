package configuration

import (
	"dolittle.io/cookie-oidc-client/completion"
	"dolittle.io/cookie-oidc-client/configuration/changes"
	"dolittle.io/cookie-oidc-client/cookies"
	"dolittle.io/cookie-oidc-client/initiation"
	"dolittle.io/cookie-oidc-client/openid"
	"dolittle.io/cookie-oidc-client/server"
	"dolittle.io/cookie-oidc-client/server/public"
	"dolittle.io/cookie-oidc-client/sessions"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

type Container struct {
	Notifier changes.ConfigurationChangeNotifier

	SessionsCreator sessions.Creator
	SessionsReader  sessions.Reader
	SessionsWriter  sessions.Writer

	CookiesWriter cookies.Writer

	OpenidInitiator openid.AuthenticationInitiator
	OpenidExchanger openid.TokenExchanger

	InitiationParser    initiation.Parser
	InitiationValidator initiation.Validator
	InitiationInitiator initiation.Initiatior

	CompletionParser    completion.Parser
	CompletionValidator completion.Validator
	CompletionCompleter completion.Completer

	CompleteHandler public.CompleteHandler
	InitiateHandler public.InitiateHandler
	Server          server.Server
}

func NewContainer(config Configuration) *Container {
	logger, _ := zap.NewDevelopment()
	container := Container{}

	container.Notifier = changes.NewConfigurationChangeNotifier(logger)

	sessionStore := gorilla.NewCookieStore([]byte("super-secret-value")) // TODO: incorporate into config so that keys can be hot-reloaded

	container.SessionsCreator = sessions.NewCreator(
		nonces.NewGenerator(config.Sessions().Nonce(), logger),
		logger)
	container.SessionsReader = sessions.NewReader(
		config.Sessions(),
		sessionStore,
		logger)
	container.SessionsWriter = sessions.NewWriter(
		config.Sessions(),
		sessionStore,
		logger)

	container.CookiesWriter = cookies.NewWriter()

	container.OpenidInitiator = openid.NewAuthenticationInitiator()

	container.InitiationParser = initiation.NewParser(
		config.Initiation(),
		logger)
	container.InitiationValidator = initiation.NewValidator(
		config.Initiation(),
		logger)
	container.InitiationInitiator = initiation.NewInitiator(
		container.InitiationValidator,
		container.SessionsCreator,
		container.OpenidInitiator,
		logger)

	container.CompletionParser = completion.NewParser(
		logger)
	container.CompletionValidator = completion.NewValidator(
		logger)
	container.CompletionCompleter = completion.NewCompleter(
		container.CompletionValidator,
		container.OpenidExchanger,
		logger)

	container.InitiateHandler = public.NewInitiateHandler(
		container.InitiationParser,
		container.InitiationInitiator,
		container.SessionsWriter)
	container.CompleteHandler = public.NewCompleteHandler(
		container.CompletionParser,
		container.SessionsReader,
		container.CompletionCompleter,
		container.CookiesWriter)
	container.Server = server.NewServer(
		config.Server(),
		container.Notifier,
		container.InitiateHandler,
		container.CompleteHandler,
		logger)

	return &container
}

// import (
// 	"context"
// 	"log"

// 	"dolittle.io/cookie-oidc-client/configuration"
// 	handlers "dolittle.io/cookie-oidc-client/handlers"
// 	"dolittle.io/cookie-oidc-client/urls"
// 	handlersUrls "dolittle.io/cookie-oidc-client/urls/handlers"
// 	"dolittle.io/cookie-oidc-client/utils"
// 	oidc "github.com/coreos/go-oidc"
// 	"github.com/gorilla/sessions"
// 	"golang.org/x/oauth2"
// )

// type Container struct {
// 	config      *configuration.Configuration
// 	oauthConfig oauth2.Config

// 	Server *Server

// 	TokenGetter            handlers.TokenGetter
// 	CookieFactory          handlers.CookieFactory
// 	NonceGenerator         utils.NonceGenerator
// 	SessionThingy          handlers.SessionThingy
// 	CallbackRedirectGetter handlersUrls.CallbackRedirectGetter
// 	ConsentPageGetter      handlersUrls.ConsentPageGetter
// 	ReturnToGetter         handlersUrls.ReturnToGetter
// }

// func getConfiguration() *configuration.Configuration {
// 	defaultConfig := configuration.GetDefaults()
// 	err := configuration.Setup(&defaultConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	config, err := configuration.Read()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return config
// }

// func createOauthConfig(config *configuration.Configuration, issuerGetter urls.OIDCProviderIssuerGetter) (oauth2.Config, error) {
// 	ctx := context.Background()
// 	issuerUrl, err := issuerGetter.GetOIDCProviderIssuerURL()
// 	if err != nil {
// 		return oauth2.Config{}, err
// 	}

// 	provider, err := oidc.NewProvider(ctx, issuerUrl.String())
// 	if err != nil {
// 		return oauth2.Config{}, err
// 	}

// 	return oauth2.Config{
// 		ClientID:     config.Provider.ClientID,
// 		ClientSecret: config.Provider.ClientSecret,
// 		Endpoint:     provider.Endpoint(),
// 		RedirectURL:  config.Provider.RedirectURL,
// 		Scopes:       config.Provider.Scopes,
// 	}, nil
// }

// func (self *Container) Setup() {
// 	self.config = getConfiguration()
// 	oauthConfig, err := createOauthConfig(self.config, urls.NewOIDCProviderIssuerGetter(self.config))
// 	if utils.TryLogIfErrorFatal(err) {
// 		return
// 	}
// 	self.oauthConfig = oauthConfig

// 	self.Server = NewServer(self.config)

// 	self.CallbackRedirectGetter = handlersUrls.NewCallbackRedirectGetter(self.config)
// 	self.ConsentPageGetter = handlersUrls.NewConsentPageGetter(self.oauthConfig)
// 	self.CookieFactory = handlers.NewOauth2CookieFactory(self.config)
// 	self.NonceGenerator = utils.NewNonceGenerator()
// 	returnToGetter, err := handlersUrls.NewReturnToGetter(self.config)
// 	if utils.TryLogIfErrorFatal(err) {
// 		return
// 	}
// 	self.ReturnToGetter = returnToGetter
// 	self.SessionThingy = handlers.NewSessionThingy(
// 		handlers.NewSessions(
// 			sessions.NewCookieStore([]byte("super-secret-value")), self.config))
// 	self.TokenGetter = handlers.NewTokenGetter(self.oauthConfig)

// }
