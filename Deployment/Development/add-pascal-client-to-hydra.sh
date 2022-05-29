#!/bin/bash
docker run --rm --network=development_authentication rancher/curl   \
    -s -X POST http://browser-hydra:4445/clients                    \
    -H 'Content-Type: application/json'                             \
    -H 'X-Forwarded-Proto: https'                                   \
    -d '
        {
            "client_name": "Studio Pascal",
            "client_id": "client-id",
            "client_secret": "client-secret",
            "redirect_uris": [
                "https://studio.localhost:8080/.auth/cookies/callback"
            ]
        }
    ' | jq
