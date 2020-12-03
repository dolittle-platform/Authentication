package sessions

import (
	"net/http"
	"net/url"

	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

// Writer writes sessions to HTTP requests
type Writer interface {
	// WriteTo writes the given Session with the HTTP ResponseWriter
	WriteTo(session *Session, r *http.Request, w http.ResponseWriter) error
}

// NewWriter returns a new Writer
func NewWriter(configuration Configuration, store gorilla.Store, logger *zap.Logger) Writer {
	return &writer{
		configuration: configuration,
		store:         store,
		logger:        logger,
	}
}

type writer struct {
	configuration Configuration
	store         gorilla.Store
	logger        *zap.Logger
}

func (w *writer) WriteTo(session *Session, request *http.Request, responseWriter http.ResponseWriter) error {
	cookie, err := w.store.New(request, w.configuration.Cookies().Name())
	if err != nil {
		return w.logAndReturnError("could not create session", err)
	}

	nonceValue := string(session.Nonce)
	returnToValue := url.URL(*session.ReturnTo)

	cookie.Values["nonce"] = nonceValue
	cookie.Values["returnTo"] = returnToValue.String()

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
