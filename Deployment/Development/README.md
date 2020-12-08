## Start with
First set the IP of your host machine (reachable from within Docker) in the `.env` file.

Append to your `/etc/hosts` file:
```
127.0.0.1 studio.localhost
127.0.0.1 local-oidc-provider
```

Then run:
```shell
docker-compose up
```

Then run:
```shell
./add-pascal-client-to-hydra.sh
```

Then run:
```shell
go run ingress.go
```

## Clean up with
```shell
docker-compose down -v
```
