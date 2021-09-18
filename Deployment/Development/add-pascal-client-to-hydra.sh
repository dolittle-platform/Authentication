#!/bin/bash
docker exec development_browser-hydra_1 hydra       \
    --endpoint http://localhost:4445 clients create             \
    --id client-id --secret client-secret                       \
    -c https://studio.localhost:8080/.auth/cookies/callback
