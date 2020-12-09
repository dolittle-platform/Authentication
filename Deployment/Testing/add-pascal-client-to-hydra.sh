#!/bin/bash
kubectl -n system-auth exec $(kubectl get pod -n system-auth -l "component=browser" -o name) -c hydra --    \
    hydra --endpoint http://localhost:4445 clients create                                                   \
    --id client-id --secret client-secret                                                                   \
    -c http://studio.localhost:8080/.auth/cookies/callback
