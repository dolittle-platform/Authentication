package openid

import "dolittle.io/cookie-oidc-client/sessions/nonces"

type AuthenticationRedirectURL string

type AuthenticationInitiator interface {
	GetAuthenticationRedirect(nonce nonces.Nonce) AuthenticationRedirectURL
}

func NewAuthenticationInitiator() AuthenticationInitiator {
	return &initiator{}
}

type initiator struct{}

func (*initiator) GetAuthenticationRedirect(nonce nonces.Nonce) AuthenticationRedirectURL {
	return "http://localhost:8080/auth"
}
