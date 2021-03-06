# [1.7.0] - 2021-7-1 [PR: #20](https://github.com/dolittle-platform/Authentication/pull/20)
## Summary

This PR adds support for using ID tokens as the value to store in the cookie. It introduces a configuration variable in `openid.token_type` that can be either `access_token` or `id_token` (defaults to `access_token` with a warning if not specified. 

Also, it seems like Azure B2C has a slightly non-standard way of acting as an OpenID Connect issuer - the well-known document requires a _flow_ query-parameter set for it to work. To support this - any query parameters set in the `openid.issuer` configuration - will be stripped away when passed to the underlying `go-oidc` library - but then added back at the end while resolving the well-known discovery document.

### Added

- Ability to use ID tokens as values in cookies. NOTE: the issuer still has to return an `access_token` in the response from the Token-endpoint for it to work.

### Fixed

- Special case handling of query-parameters in the configured issuer URL to support Azure B2C.


# [1.6.1] - 2021-3-24 [PR: #19](https://github.com/dolittle-platform/Authentication/pull/19)
Nothing to see here really, just cleaning up my mess.


