## Start with
First set the IP of your host machine (reachable from within Docker) in the `.env` file.

Append to your `/etc/hosts` file:
```
127.0.0.1 local.dolittle.studio
```

Then run:
```shell
docker-compose up
```

Then run:
```shell
go run ingress.go
```

## Clean up with
```shell
docker-compose down
```