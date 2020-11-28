package handlers

import (
	"context"
	"log"
	"net/http"

	urls "dolittle.io/cookie-oidc-client/urls/handlers"
	"dolittle.io/cookie-oidc-client/utils"
	"golang.org/x/oauth2"
)

type TokenGetter interface {
	Get(ctx context.Context, r *http.Request) (*oauth2.Token, error)
}

type tokenGetter struct {
	oauthConfig oauth2.Config
}

func NewTokenGetter(oauthConfig oauth2.Config) *tokenGetter {
	return &tokenGetter{oauthConfig}
}

func (self *tokenGetter) Get(ctx context.Context, r *http.Request) (*oauth2.Token, error) {
	return self.oauthConfig.Exchange(ctx, r.URL.Query().Get("code"))
}

type callbackHandler struct {
	callbackRedirectGetter urls.CallbackRedirectGetter
	cookieFactory          CookieFactory
	tokenGetter            TokenGetter
}

func NewCallbackHandler(
	callbackRedirectGetter urls.CallbackRedirectGetter,
	cookieFactory CookieFactory,
	tokenGetter TokenGetter) http.Handler {
	return &callbackHandler{
		callbackRedirectGetter,
		cookieFactory,
		tokenGetter,
	}
}

func (self *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received callback")

	ctx := context.Background()
	oauth2Token, err := self.tokenGetter.Get(ctx, r)
	if utils.TryLogIfErrorHttp(err, w) {
		return
	}

	// Do more stuff to contents
	cookie := self.cookieFactory.Create(oauth2Token.AccessToken)
	log.Println("Got access token, now setting cookie")
	http.SetCookie(w, cookie)

	callbackURL, err := self.callbackRedirectGetter.GetCallbackRedirectURL()

	if utils.TryLogIfErrorHttp(err, w) {
		return
	}
	callbackURLString := callbackURL.String()
	log.Println("Redirecting to ", callbackURLString)

	http.Redirect(w, r, callbackURLString, http.StatusFound)
}
