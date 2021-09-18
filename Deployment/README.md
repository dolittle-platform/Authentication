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
> NEW NOTES:
> Chrome allow localhost invalid certs

> The files and scripts referenced in this section is in the `/Deployment/Development` directory in this repository.
> The setup relies on a self-signed SSL certificate, so you need to accept that for it to work.
> If you're using Chrome, you need to enable insecure sertificates on localhost by going to [chrome://flags/#allow-insecure-localhost](chrome://flags/#allow-insecure-localhost).

##### Starting up
To build and start up the current code, run:
```shell
docker-compose up -d
```

> It takes a litle while to boot up, so give it a few seconds.
> You can run `docker logs development_browser-pascal_1 -f` and wait for the message "OpenID Connect issuer ready" to appear.

Then to add configuration to Hydra, run:
```shell
./add-pascal-client-to-hydra.sh
```

##### Testing it out
Navigate to https://studio.localhost:8080/, and log in in with email `do@do.do`, and password `password`.

The first time logging in from a fresh setup, the user will not be assigned to any tenants. To add a tenant to a user run:
```shell
./add-tenant-to-kratos-identity.sh <email> <tenant-id>
```
> The default config is set up with with a tenant mapping for `dolittle` and `tenant-a`, so try out:
```shell
./add-tenant-to-kratos-identity.sh do@do.do dolittle
```

Refresh the select tenant page if you just added another tenant, and select the tenant.

You should then be presented with the amazing Dolittle spinner page - congratulations!

##### Tearing down
Shut down the containers and the ingress, and run:
```shell
docker-compose down -v
```
This will clean up everything created and clear out databases, so you'll need to add tenants to users again the next time.

## 3. Running in a local Kubernetes cluster
