#!/bin/bash
docker exec dolittle_authentication_browser-hydra_1 hydra       \
    --endpoint http://localhost:4445 clients create             \
    --id client-id --secret client-secret                       \
    -c http://studio.localhost:8080/.auth/cookies/callback
