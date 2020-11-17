# Authentication

## Running locally
...

kubectl -n system-auth port-forward postgresql-0 8080:80


## Paths

`/spa/login`


/ -> apiserver-proxy:80/
/.ory/kratos/public -> kratos:4433/
/oauth2 -> hydra:4444/oauth2
/.well-known -> hydra:4444/well-known

k8s.dolittle.studio/ -> "k8 apiserver proxy path"

/ -> studio
/.auth/select-tenant -> "select tenant page"
/.auth/login -> "select login provider page"

/.auth/initiate -> "cookie thing"
/.auth/callback -> "cookie thing"

/.openid -> hydra:public/

/.ory/kratos/public -> kratos:public/
