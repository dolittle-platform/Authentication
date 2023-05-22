# [3.1.0] - 2023-5-22 [PR: #34](https://github.com/dolittle-platform/Authentication/pull/34)
## Summary

Introduce changes to Authentication that give new users a slightly better experience with Studio by presenting them with a "no-tenant" landing page with links to other Dolittle resources. The main intent is for them not to "be left hanging". This is a first step in the direction of self-serviced sign-ups.


### Added

- [Login/Backend]: Add a rule in the authentication backend to redirect a user with no tenants to a new "no-tenant" page.
- [Login/Web]: Add a link on the select customer page to "Get access" that takes the user to the "no-tenant" page
- [Login/Web]: Add a new "no-tenant" page that serves as a landing page for new users of Studio and has links to further learning 

### Changed

- [Login/Web]: Rebranded from Dolittle to Aigonix


# [3.0.1] - 2023-3-13 [PR: #33](https://github.com/dolittle-platform/Authentication/pull/33)
## Summary

Investigated how to redirect a user when only one Tenant is available. By doing this, I cleaned up the code a bit.


# [3.0.0] - 2022-10-19 [PR: #32](https://github.com/dolittle-platform/Authentication/pull/32)
## Summary

Adds support for running a single instance of Pascal to serve multiple hosts. The configuration now expects a list of allowed hostnames (`host[:port]`) that should be served (any other host is rejected with 404), and every other URL is constructed on the fly based on the host of the handled request. Every other URL in the configuration is now expected to be absolute relative to the current host.

Example config:
```
serve:
  port: 80

  hosts:
    - localhost:80

  paths:
    initiate: /initiate
    complete: /callback
    logout: /logout

urls:
  error: /error
  return:
    query_parameter: return_to
    default:
      login: /return
      logout: /return
    allowed:
      - /
    mode: strict

sessions:
  nonce_length: 80
  lifetime: 5m
  cookies:
    name: .dolittle.pascal.session
    secure: true
    samesite: lax
    path: /
  keys:
    - hash: KEY-USED-TO-SIGN-SESSION-COOKIES-SHOULD-BE-64-BYTES-LONG--------
      block: ENCRYPTION-KEY-SHOULD-BE-32-BYTS

openid:
  issuer: http://localhost:80
  client:
    id: client-id
    secret: client-secret
  scopes:
    - openid
  token_type: access_token
  redirect: /callback

cookies:
  name: .dolittle.pascal.login
  secure: true
  samesite: lax
  path: /
```


### Added

- [Pascal]: can handle requests for multiple hosts/domains.

### Changed

- [Pascal]: requires config of allowed hosts to serve, and redirect URLs now need to be just absolute paths without a host.
- [Pascal]: skip logging stack trace at warning log messages.

### Fixed

- [Login/Web]: upgrade `rest-hooks` packages so that the build un-breaks.


# [2.5.2] - 2022-8-25 [PR: #31](https://github.com/dolittle-platform/Authentication/pull/31)
## Summary

A few small UX performance improvements for the Login service - to make the page load faster and look better while loading for the end-user.

### Added

- [Login/Web] Background colour to the `<html>` tag to avoid flashing a white page while assets are loading
- [Login/Web] Add `font-display: swap` to make text display while loading fonts
- [Login/Backend] Add `Cache-Control` headers to served content to improve caching loading speed.


# [2.5.1] - 2022-8-25 [PR: #30](https://github.com/dolittle-platform/Authentication/pull/30)
## Summary

Changed background Dolittle logo color for having more contrast.

<img width="1285" alt="Screenshot 2022-08-18 at 11 48 22" src="https://user-images.githubusercontent.com/19160439/185352908-6a58bf4d-242f-4075-aacf-21e8a40b57b4.png">

### Changed

- Background logo color


# [2.5.0] - 2022-8-18 [PR: #29](https://github.com/dolittle-platform/Authentication/pull/29)
## Summary

Fixes a problem where it would not be possible to get out of some error states. The issue was that the error page asks you to log out and try again, but if the user is not logged in - the logout flow would fail and send you back to the error page. Thus creating an infinite loop.

### Added

- [Login/Backend] Configuration for redirect URL when logging out and there is no logged in user.

### Fixed

- [Login/Backend] If there is no session in Kratos (initiate logout flow returns 401), redirect to the configured logged-out URL.


# [2.4.2] - 2022-8-16 [PR: #28](https://github.com/dolittle-platform/Authentication/pull/28)
## Summary

Fixed 'Select your customer' buttons width to be as wide as the widest button is.

<img width="633" alt="Screenshot 2022-08-12 at 18 37 49" src="https://user-images.githubusercontent.com/19160439/184390398-a78006c3-185b-45e6-a568-48e8a5a68dc8.png">

### Fixed

- Buttons width


# [2.4.1] - 2022-8-11 [PR: #27](https://github.com/dolittle-platform/Authentication/pull/27)
## Summary

Some refactoring of the code to keep it consistent but still up to date with the latest changes, and bringing back the ability to configure all "Dolittle Studio" specific things out so it can be used by others.


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


