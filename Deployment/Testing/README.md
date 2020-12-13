## Start with
First set the IP of your host machine (reachable from within Docker) in the `hostAliases` field of the `browser` deployment in `browser.yml` (line ~93) file.

Append to your `/etc/hosts` file:
```
127.0.0.1 studio.localhost
127.0.0.1 local-oidc-provider
```

Build local Docker images with (in the `../Development` directory)
```shell
docker-compose build
```

Then run:
```shell
kubectl apply -f .
```

Then run:
```shell
./add-pascal-client-to-hydra.sh
```

## Clean up with
```shell
kubectl delete ns system-ingress system-auth studio
```
