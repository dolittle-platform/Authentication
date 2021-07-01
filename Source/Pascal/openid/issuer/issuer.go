package issuer

import (
	"context"
	"net/url"

	"dolittle.io/pascal/sessions/nonces"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

const idTokenKey = "id_token"

type Issuer interface {
	GetAuthenticationRedirectURL(nonce nonces.Nonce) (string, error)
	ExchangeCodeForAccessToken(code string) (*Token, error)
	ExchangeCodeForIDToken(code string) (*Token, error)
}

type issuer struct {
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
	config   *oauth2.Config
}

func NewIssuer(issuerURL *url.URL, clientId, clientSecret string, scopes []string, redirectUrl *url.URL) (Issuer, error) {
	issuerUrl, query := splitIssuerUrlAndQuery(issuerURL)

	ctx := context.Background()
	if len(query) > 0 {
		client := getHttpClientFor(issuerUrl, query)
		ctx = context.WithValue(ctx, oauth2.HTTPClient, client)
	}

	provider, err := oidc.NewProvider(ctx, issuerUrl.String())
	if err != nil {
		return nil, err
	}

	config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       getScopesWithOpenID(scopes),
		RedirectURL:  redirectUrl.String(),
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: clientId,
	})

	return &issuer{
		provider: provider,
		verifier: verifier,
		config:   config,
	}, nil
}

func (i *issuer) GetAuthenticationRedirectURL(nonce nonces.Nonce) (string, error) {
	return i.config.AuthCodeURL(string(nonce)), nil
}

func (i *issuer) ExchangeCodeForAccessToken(code string) (*Token, error) {
	token, err := i.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return &Token{
		Value:   token.AccessToken,
		Expires: token.Expiry,
	}, nil
}

func (i *issuer) ExchangeCodeForIDToken(code string) (*Token, error) {
	token, err := i.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	idTokenValue, ok := token.Extra(idTokenKey).(string)
	if !ok {
		return nil, ErrIDTokenMissing
	}

	idToken, err := i.verifier.Verify(context.Background(), idTokenValue)
	if err != nil {
		return nil, err
	}

	return &Token{
		Value:   idTokenValue,
		Expires: idToken.Expiry,
	}, nil
}

func splitIssuerUrlAndQuery(issuerUrl *url.URL) (*url.URL, url.Values) {
	if len(issuerUrl.Query()) > 0 {
		return &url.URL{
			Scheme:  issuerUrl.Scheme,
			Opaque:  issuerUrl.Opaque,
			User:    issuerUrl.User,
			Host:    issuerUrl.Host,
			Path:    issuerUrl.Path,
			RawPath: issuerUrl.RawPath,
		}, issuerUrl.Query()
	} else {
		return issuerUrl, nil
	}
}

func getScopesWithOpenID(scopes []string) []string {
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			return scopes
		}
	}
	return append(scopes, oidc.ScopeOpenID)
}
