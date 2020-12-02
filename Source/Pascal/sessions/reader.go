package sessions

import (
	"net/http"
	"net/url"

	"dolittle.io/pascal/sessions/nonces"
	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

// Reader reads sessions from HTTP requests
type Reader interface {
	// ReadFrom returns the Session stored in HTTP request
	ReadFrom(r *http.Request) (*Session, error)
}

// NewReader returns a new Reader
func NewReader(configuration Configuration, store gorilla.Store, logger *zap.Logger) Reader {
	return &reader{
		configuration: configuration,
		store:         store,
		logger:        logger,
	}
}

type reader struct {
	configuration Configuration
	store         gorilla.Store
	logger        *zap.Logger
}

func (r *reader) ReadFrom(request *http.Request) (*Session, error) {
	cookie, err := r.store.Get(request, r.configuration.Cookies().Name())
	if err != nil {
		return r.logAndReturnError("could not read session from request", err)
	}

	nonceValue, hasNonce := cookie.Values["nonce"]
	if !hasNonce {
		return r.logAndReturnError("session in request is incomplete", ErrSessionIsMissingNonce)
	}
	nonceString, nonceIsString := nonceValue.(string)
	if !nonceIsString {
		return r.logAndReturnError("session in request is invalid", ErrNonceWasNotAString)
	}

	returnToValue, hasReturnTo := cookie.Values["returnTo"]
	if !hasReturnTo {
		return r.logAndReturnError("session in request is incomplete", ErrSessionIsMissingReturnTo)
	}
	returnToString, returnToIsString := returnToValue.(string)
	if !returnToIsString {
		return r.logAndReturnError("session in request is invalid", ErrReturnToWasNotAURL)
	}
	returnTo, err := url.Parse(returnToString)
	if err != nil {
		r.logger.Error("could not parse session return to as URL", zap.Error(err))
		return r.logAndReturnError("session in request is invalid", ErrReturnToWasNotAURL)
	}

	return &Session{
		Nonce:    nonces.Nonce(nonceString),
		ReturnTo: returnTo,
	}, nil
}

func (r *reader) logAndReturnError(message string, err error) (*Session, error) {
	r.logger.Error(message, zap.Error(err))
	return nil, err
}
