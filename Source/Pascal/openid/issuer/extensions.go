package issuer

import (
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"net/url"
	"strings"
)

type issuerExtensions struct {
	endpointAuthStyle  oauth2.AuthStyle
	supportsRevocation bool
	revocationEndpoint string
	supportsLogout     bool
	logoutEndpoint     string
}

func getIssuerExtensionsFrom(provider *oidc.Provider) (*issuerExtensions, error) {
	extra := &extraClaims{}
	if err := provider.Claims(extra); err != nil {
		return nil, err
	}

	extensions := &issuerExtensions{}

	if extra.authModeIsSupported("client_secret_basic") {
		extensions.endpointAuthStyle = oauth2.AuthStyleInHeader
	} else if extra.authModeIsSupported("client_secret_post") {
		extensions.endpointAuthStyle = oauth2.AuthStyleInParams
	} else {
		return nil, ErrUnknownIssuerTokenEndpointAuthMethod
	}

	if _, err := url.ParseRequestURI(extra.RevocationEndpoint); err == nil {
		extensions.supportsRevocation = true
		extensions.revocationEndpoint = extra.RevocationEndpoint
	}

	if _, err := url.ParseRequestURI(extra.EndSessionEndpoint); err == nil {
		extensions.supportsLogout = true
		extensions.logoutEndpoint = extra.EndSessionEndpoint
	}

	return extensions, nil
}

type extraClaims struct {
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
}

func (c *extraClaims) authModeIsSupported(authMethod string) bool {
	for _, supportedMethod := range c.TokenEndpointAuthMethodsSupported {
		if strings.EqualFold(authMethod, supportedMethod) {
			return true
		}
	}
	return false
}
