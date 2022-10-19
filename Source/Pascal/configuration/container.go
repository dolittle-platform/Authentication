package configuration

import (
	"dolittle.io/pascal/completion"
	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/initiation"
	"dolittle.io/pascal/logout"
	"dolittle.io/pascal/openid"
	"dolittle.io/pascal/server"
	"dolittle.io/pascal/server/public"
	"dolittle.io/pascal/sessions"
	"dolittle.io/pascal/sessions/nonces"
	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

type Container struct {
	Notifier changes.ConfigurationChangeNotifier

	SessionStore      gorilla.Store
	SessionsCreator   sessions.Creator
	SessionsDestroyer sessions.Destroyer
	SessionsReader    sessions.Reader
	SessionsWriter    sessions.Writer

	CookiesWriter   cookies.Writer
	CookiesReader   cookies.Reader
	CookiesCrumbler cookies.Crumbler

	OpenidInitiator openid.AuthenticationInitiator
	OpenidExchanger openid.TokenExchanger
	OpenidRevoker   openid.TokenRevoker

	InitiationParser    initiation.Parser
	InitiationValidator initiation.Validator
	InitiationInitiator initiation.Initiatior

	CompletionParser    completion.Parser
	CompletionValidator completion.Validator
	CompletionCompleter completion.Completer

	LogoutParser    logout.Parser
	LogoutInitiator logout.Initiator

	CompleteHandler public.CompleteHandler
	InitiateHandler public.InitiateHandler
	LogoutHandler   public.LogoutHandler
	Server          server.Server
}

func NewContainer(config Configuration) (*Container, error) {
	logger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.ErrorLevel))
	container := Container{}

	container.Notifier = changes.NewConfigurationChangeNotifier(logger)
	config.OnChange(container.Notifier.TriggerChanged)

	cookieStore, err := sessions.NewCookieStore(
		config.Sessions(),
		container.Notifier)
	if err != nil {
		return nil, err
	}
	container.SessionStore = cookieStore

	container.SessionsCreator = sessions.NewCreator(
		nonces.NewGenerator(
			config.Sessions().Nonce(),
			logger),
		logger)
	container.SessionsDestroyer = sessions.NewDestroyer(
		config.Sessions(),
		container.SessionStore,
		logger)
	container.SessionsReader = sessions.NewReader(
		config.Sessions(),
		container.SessionStore,
		logger)
	container.SessionsWriter = sessions.NewWriter(
		config.Sessions(),
		container.SessionStore,
		logger)

	container.CookiesWriter = cookies.NewWriter(
		config.Cookies())
	container.CookiesReader = cookies.NewReader(
		config.Cookies())
	container.CookiesCrumbler = cookies.NewCrumbler(
		config.Cookies())

	initiator, err := openid.NewAuthenticationInitiator(
		config.OpenID(),
		container.Notifier,
		logger)
	if err != nil {
		return nil, err
	}
	container.OpenidInitiator = initiator
	exchanger, err := openid.NewTokenExchanger(
		config.OpenID(),
		container.Notifier,
		logger)
	if err != nil {
		return nil, err
	}
	container.OpenidExchanger = exchanger
	revoker, err := openid.NewTokenRevoker(
		config.OpenID(),
		container.Notifier,
		logger)
	if err != nil {
		return nil, err
	}
	container.OpenidRevoker = revoker

	container.InitiationParser = initiation.NewParser(
		config.Redirects(),
		logger)
	container.InitiationValidator = initiation.NewValidator(
		config.Redirects(),
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

	container.LogoutParser = logout.NewParser(
		config.Redirects(),
		container.CookiesReader,
		logger)
	container.LogoutInitiator = logout.NewInitiator(
		container.OpenidRevoker,
		container.OpenidInitiator,
		logger)

	container.InitiateHandler = public.NewInitiateHandler(
		container.InitiationParser,
		container.InitiationInitiator,
		container.SessionsWriter)
	container.CompleteHandler = public.NewCompleteHandler(
		container.CompletionParser,
		container.SessionsReader,
		container.SessionsDestroyer,
		container.CompletionCompleter,
		container.CookiesWriter)
	container.LogoutHandler = public.NewLogoutHandler(
		container.CookiesCrumbler,
		container.LogoutParser,
		container.LogoutInitiator,
		logger)
	container.Server = server.NewServer(
		config.Server(),
		container.Notifier,
		container.InitiateHandler,
		container.CompleteHandler,
		container.LogoutHandler,
		logger)

	return &container, nil
}
