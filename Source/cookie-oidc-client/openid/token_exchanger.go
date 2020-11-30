package openid

import "golang.org/x/oauth2"

type AuthenticationCode string

type TokenExchanger interface {
	Exchange(code AuthenticationCode) (*oauth2.Token, error)
}
