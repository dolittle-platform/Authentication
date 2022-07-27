# [2.4.0] - 2022-7-27 [PR: #26](https://github.com/dolittle-platform/Authentication/pull/26)
## Summary

Improved login flow with added design and changed content.

### Changed

- Theme updated


# [2.3.0] - 2022-7-15 [PR: #25](https://github.com/dolittle-platform/Authentication/pull/25)
## Summary

Added favicon in ico format.

### Added

- favicon.ico file


# [2.2.0] - 2022-6-16 [PR: #24](https://github.com/dolittle-platform/Authentication/pull/24)
## Summary

Introduces a new `dolittle/login-provider-icons` Docker image. This is an NGINX server which serves baked Login Provider icons to be used in the Login Web frontend. This way we have a little more control of the icon URLs, and we don't have to go cross-domain to fetch the images.

### Added

- A new Login Provider Icons NGINX server that can be used to serve the correct provider icons for the Login Web frontend.


# [2.1.0] - 2022-6-15 [PR: #23](https://github.com/dolittle-platform/Authentication/pull/23)
## Summary

Implements even newer design for the Login frontend. Simplified dark theme and more responsive.


# [2.0.1] - 2022-5-29 [PR: #22](https://github.com/dolittle-platform/Authentication/pull/22)
## Summary

Fixes Docker build of _Login_.

### Fixed

- A typo in the _Login_ `Dockerfile` that caused the previous release to fail.


# [2.0.0] - 2022-5-29 [PR: #21](https://github.com/dolittle-platform/Authentication/pull/21)
## Summary

A rather large rewrite of multiple components in the authentication system to bring:
- New design and UX improvements (thanks Liz!)
- The ability to log out properly
- Make the UI configurable at runtime to make it easier to whitelabel

### Added

- Configuration of the Login frontend at runtime
- Pascal and Login now supports logging out (ending the users browser session and revoking tokens properly)

### Changed

- New design and UX in the Login frontend
- Upgraded to React 18 and MUI 5
- Upgraded Hydra, Oathkeeper and Kratos to newer versions
- Hardening of cookie configuration, and enabling Hydra to run with terminated TLS (see Docker-Compose sample setup)
- New Docker-Compose local development setup that requires less hacks to get going


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


