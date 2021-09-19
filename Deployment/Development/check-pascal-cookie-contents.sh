#!/bin/bash
docker run --rm --network=development_authentication rancher/curl   \
    -s -X POST http://browser-hydra:4445/oauth2/introspect          \
    -H 'Content-Type: application/x-www-form-urlencoded'            \
    -H 'X-Forwarded-Proto: https'                                   \
    --data-urlencode "token=$1" | jq
