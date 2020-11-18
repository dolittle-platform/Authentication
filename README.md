# Authentication

## Running locally
...

```shell
kubectl -n system-auth port-forward <postgres-pod> 8080:80
```

```shell
kubectl -n system-auth $(kubectl get pod -l "component=hydra" -o name -n system-auth) exec -- hydra --endpoint http://localhost:4445 clients create --id do --secret little -c http://localhost:8080/.auth/callback/
kubectl -n system-auth exec $(kubectl get pod -l "component=hydra" -o name -n system-auth) -- hydra --endpoint http://localhost:4445 clients list
```

Add to your /etc/hosts (bottom is a good idea)
```
127.0.0.1 oidc-provider.oidc-provider.svc.cluster.local
```

To get users in kratos
```shell
kubectl -n system-auth exec $(kubectl get pod -l "component=kratos" -o name -n system-auth) -- kratos --endpoint=http://localhost:4434 identities list -f=json-pretty
```

To add tenants to a user in kratos

```shell
kubectl -n system-auth port-forward $(kubectl get pod -l "component=kratos" -o name -n system-auth) 4434:4434

curl -X PUT http://localhost:4434/identities/{id} \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' --data @- <<EOF
    {
        "schema_id": "default",
        "traits": {
            "email": "{email}",
            "tenants": [ "tenant-a", "tenant-b" ]
        }
    }
EOF
```

```shell
curl -X PUT http://localhost:4434/identities/0beb6a1f-7baf-4990-aa32-2aa40989281f \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' --data @- <<EOF
    {
        "schema_id": "default",
        "traits": {
            "email": "do@do.do",
            "tenants": [ "tenant-a", "tenant-b" ]
        }
    }
EOF
```

## Paths


```
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
```