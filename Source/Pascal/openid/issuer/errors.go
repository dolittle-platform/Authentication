package issuer

import "errors"

var (
	ErrIDTokenMissing                       = errors.New("ID token not received in token from issuer")
	ErrUnknownIssuerTokenEndpointAuthMethod = errors.New("the issuers token_endpoint_auth_methods_supported does not contain any known auth methods")
	ErrRevocationIsNotSupported             = errors.New("the issuer does not support token revocation")
	ErrTokenRevokeReturnedNotOK             = errors.New("the issuer did not return 200 while revoking a token")
)
