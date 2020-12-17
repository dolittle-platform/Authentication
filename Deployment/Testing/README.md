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

Wait for a little while until the databases are initialised (it's ready when all pods are running `kubectl -n system-auth get pod`), then run:
```shell
./add-pascal-client-to-hydra.sh
```

Then you can proceed to `http://studio.localhost:8080` and follow the links through the login.
After logging in, you should get to the _Select Tenant_ page, with no options.

Find the ID of your user with:
```shell
./get-kratos-identities.sh
```

And add a tenant to your user:
```shell
./add-tenant-to-kratos-identity.sh <user-id-from-previous-output> tenant-a
```
You can add more tenants with the same command.
If you only have one tenant, you will not be presented with the page to select tenant.

## Clean up with
```shell
kubectl delete ns system-ingress system-auth studio
```
