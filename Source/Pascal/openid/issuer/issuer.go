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
	GetAuthenticationRedirectURL(host string, nonce nonces.Nonce) (string, error)
	ExchangeCodeForAccessToken(code string) (*Token, error)
	ExchangeCodeForIDToken(code string) (*Token, error)
	RevocationIsSupported() bool
	RevokeToken(*Token) error
	LogoutIsSupported() bool
	GetLogoutRedirectURL(idTokenHint, state, returnTo string) (string, error)
}

type issuer struct {
	provider   *oidc.Provider
	verifier   *oidc.IDTokenVerifier
	config     *oauth2.Config
	extensions *issuerExtensions
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

	extensions, err := getIssuerExtensionsFrom(provider)
	if err != nil {
		return nil, err
	}

	return &issuer{
		provider:   provider,
		verifier:   verifier,
		config:     config,
		extensions: extensions,
	}, nil
}

func (i *issuer) GetAuthenticationRedirectURL(host string, nonce nonces.Nonce) (string, error) {
	redirectConfig := &oauth2.Config{
		ClientID:     i.config.ClientID,
		ClientSecret: i.config.ClientSecret,
		Endpoint:     i.config.Endpoint,
		Scopes:       i.config.Scopes,
		RedirectURL:  host + i.config.RedirectURL,
	}

	return redirectConfig.AuthCodeURL(string(nonce)), nil
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

func (i *issuer) RevocationIsSupported() bool {
	return i.extensions.supportsRevocation
}

func (i *issuer) RevokeToken(token *Token) error {
	if !i.RevocationIsSupported() {
		return ErrRevocationIsNotSupported
	}

	return revokeToken(token.Value, i.extensions.revocationEndpoint, i.config.ClientID, i.config.ClientSecret, i.extensions.endpointAuthStyle)
}

func (i *issuer) LogoutIsSupported() bool {
	return i.extensions.supportsLogout
}

func (i *issuer) GetLogoutRedirectURL(idTokenHint, state, returnTo string) (string, error) {
	if !i.LogoutIsSupported() {
		return "", ErrLogoutIsNotSupported
	}

	v := url.Values{}
	if idTokenHint != "" {
		v.Set("id_token_hint", idTokenHint)
	}
	if state != "" {
		v.Set("state", state)
	}
	if returnTo != "" {
		v.Set("post_logout_uri", returnTo)
	}

	if len(v) == 0 {
		return i.extensions.logoutEndpoint, nil
	}

	return i.extensions.logoutEndpoint + "?" + v.Encode(), nil

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
