package main

import (
	"net/http"

	"dolittle.io/cookie-oidc-client/handlers"
)

func main() {
	container := Container{}
	container.Setup()

	initiateHandler := handlers.NewInitiateHandler(
		container.NonceGenerator,
		container.SessionThingy,
		container.ReturnToGetter,
		container.ConsentPageGetter)
	http.Handle("/initiate/", initiateHandler)
	callbackHandler := handlers.NewCallbackHandler(
		container.CallbackRedirectGetter,
		container.CookieFactory,
		container.TokenGetter)
	http.Handle("/callback/", callbackHandler)

	container.Server.Start()
}
