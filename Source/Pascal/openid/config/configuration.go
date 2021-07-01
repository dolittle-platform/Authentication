package config

import "net/url"

type TokenType string

const (
	AccessToken TokenType = "access_token"
	IDToken     TokenType = "id_token"
)

type Configuration interface {
	Issuer() *url.URL
	ClientID() string
	ClientSecret() string
	Scopes() []string
	TokenType() TokenType
	RedirectURL() *url.URL
}
