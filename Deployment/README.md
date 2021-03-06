# How to run locally
There are three ways to run the code in this repository locally:
1. Mocking out the whole backend for just working on the frontend
2. Running with Docker Compose (the setup is a bit simpler than the whole Kubernetes environment)
3. Running in a local Kubernetes cluster to most accurately simulate the runtime environment.

## 1. Running with a mocked backend
> The files and scripts referenced in this section is in the `/Deployment/Frontend` directory in this repository.

Install and run the express server:
```shell
yarn
yarn start
```

Start the frontend code in `Source/Login/Web`, navigate to http://localhost:8080.

Select a provider to log in, then a tenant, and you should be presented with a page saying you are logged in. The page also has a link to restart the login process, or to view the error page in the frontend.

To test out with different sets of providers or tenants, you can modify the `identityProviders` and `availableTenants` variables on the top of the `index.js` file.

> Note: no special handling is required for a case where a single tenant is available for a user, this is handled by the backend which would jump over the select tenant page.

## 2. Running with Docker Compose
> The files and scripts referenced in this section is in the `/Deployment/Development` directory in this repository.
> It also requires a local installation of Go.

##### Configuring the network
First, set the IP of your host machine (reachable from a Docker container) to the `HOST_IP` in the `.env` file.

> If you're on a Linux host, you can find the IP from the `inet` field of the `docker0` interface using `ip addr show docker0`

Then, append to your `/etc/hosts` file the following:
```
127.0.0.1 studio.localhost
127.0.0.1 local-oidc-provider
```

##### Starting up
To build and start up the current code, run:
```shell
docker-compose up
```

In another terminal, run:
```shell
./add-pascal-client-to-hydra.sh
```
And then lastly, run:
```shell
go run ingress.go
```

##### Testing it out
Navigate to http://studio.localhost:8080/, and log in in with email `do@do.do`, and password `password`.

The first time logging in from a fresh setup, the user will not be assigned to any tenants. To add a tenant to a user run:
```shell
./add-tenant-to-kratos-identity.sh <email> <tenant-id>
```
> The default config is set up with with a tenant mapping for `dolittle` and `tenant-a`, so try out:
> ./add-tenant-to-kratos-identity.sh do@do.do dolittle

Refresh the select tenant page if you just added another tenant, and select the tenant.

You should then be presented with the amazing Dolittle spinner page - congratulations!

##### Tearing down
Shut down the containers and the ingress, and run:
```shell
docker-compose down
```
This will clean up everything created and clear out databases, so you'll need to add tenants to users again the next time.

## 3. Running in a local Kubernetes cluster
