package handlers

import (
	"net/http"

	urls "dolittle.io/cookie-oidc-client/urls/handlers"
	"dolittle.io/cookie-oidc-client/utils"
)

type initiateHandler struct {
	nonceGenerator    utils.NonceGenerator
	sessionThingy     SessionThingy
	returnToGetter    urls.ReturnToGetter
	consentPageGetter urls.ConsentPageGetter
}

func NewInitiateHandler(
	nonceGenerator utils.NonceGenerator,
	sessionThingy SessionThingy,
	returnToGetter urls.ReturnToGetter,
	consentPageGetter urls.ConsentPageGetter) http.Handler {
	return &initiateHandler{
		nonceGenerator,
		sessionThingy,
		returnToGetter,
		consentPageGetter,
	}
}

func (self *initiateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnTo, err := self.returnToGetter.GetReturnToURL(r)
	if utils.TryLogIfErrorHttp(err, w) {
		return
	}

	nonce, err := self.nonceGenerator.Generate()
	if utils.TryLogIfErrorHttp(err, w) {
		return
	}

	err = self.sessionThingy.SetSession(r)
	if utils.TryLogIfErrorHttp(err, w) {
		return
	}
	err = self.sessionThingy.WriteSession(nonce, returnTo, w, r)
	if utils.TryLogIfErrorHttp(err, w) {
		return
	}

	consentPageURL, err := self.consentPageGetter.GetConsentPageURL(nonce)

	http.Redirect(w, r, consentPageURL.String(), http.StatusFound)
}
