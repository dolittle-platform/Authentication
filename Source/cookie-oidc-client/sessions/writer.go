package sessions

import (
	"net/http"

	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

// Writer writes sessions to HTTP requests
type Writer interface {
	// WriteTo writes the given Session with the HTTP ResponseWriter
	WriteTo(session *Session, r *http.Request, w http.ResponseWriter) error
}

// NewWriter returns a new Writer
func NewWriter(configuration Configuration, store gorilla.Store, logger zap.Logger) Writer {
	return &writer{
		configuration: configuration,
		store:         store,
		logger:        logger,
	}
}

type writer struct {
	configuration Configuration
	store         gorilla.Store
	logger        zap.Logger
}

func (w *writer) WriteTo(session *Session, request *http.Request, responseWriter http.ResponseWriter) error {
	cookie, err := w.store.New(request, w.configuration.CookieName())
	if err != nil {
		return w.logAndReturnError("could not create session", err)
	}

	cookie.Values["nonce"] = session.Nonce
	cookie.Values["returnTo"] = session.ReturnTo

	err = w.store.Save(request, responseWriter, cookie)
	if err != nil {
		return w.logAndReturnError("could not save session", err)
	}
	return nil
}

func (w *writer) logAndReturnError(message string, err error) error {
	w.logger.Error(message, zap.Error(err))
	return err
}
