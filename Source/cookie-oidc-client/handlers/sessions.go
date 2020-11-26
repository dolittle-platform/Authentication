package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"dolittle.io/cookie-oidc-client/configuration"
	gorilla_sessions "github.com/gorilla/sessions"
)

type Sessions interface {
	Get(r *http.Request) (*gorilla_sessions.Session, error)
}
type sessions struct {
	sessionStore     gorilla_sessions.Store
	sessionStoreName string
}

var _ Sessions = new(sessions)

func NewSessions(sessionStore gorilla_sessions.Store, config *configuration.Configuration) *sessions {
	return &sessions{sessionStore, config.SessionStoreName}
}

func (self *sessions) Get(r *http.Request) (*gorilla_sessions.Session, error) {
	return self.sessionStore.New(r, self.sessionStoreName)
}

type SessionThingy interface {
	SetSession(r *http.Request) error
	WriteSession(nonce string, returnTo *url.URL, w http.ResponseWriter, r *http.Request) error
	ReadSession(w http.ResponseWriter, r *http.Request) (nonce string, returnTo *url.URL, err error)
}

type sessionThingy struct {
	sessions Sessions
	session  *gorilla_sessions.Session
}

var _ SessionThingy = new(sessionThingy)

func NewSessionThingy(sessions Sessions) *sessionThingy {
	return &sessionThingy{sessions, nil}
}

func (self *sessionThingy) SetSession(r *http.Request) error {
	session, err := self.sessions.Get(r)
	if err != nil {
		return err
	}
	self.session = session
	return nil
}

func (self *sessionThingy) WriteSession(nonce string, returnTo *url.URL, w http.ResponseWriter, r *http.Request) error {
	if self.session == nil {
		return errors.New("session is not set")
	}
	self.session.Values["nonce"] = nonce
	self.session.Values["return_to"] = returnTo.String()
	return self.session.Save(r, w)
}

func (self *sessionThingy) ReadSession(w http.ResponseWriter, r *http.Request) (nonce string, returnTo *url.URL, err error) {
	if self.session == nil {
		return "", nil, errors.New("session is not set")
	}
	nonce = self.session.Values["nonce"].(string)
	returnTo, err = url.Parse(self.session.Values["return_to"].(string))
	return
}
