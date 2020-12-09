#!/bin/bash
kubectl -n system-auth exec $(kubectl get pod -n system-auth -l "component=login" -o name) -c kratos -- \
    kratos --endpoint http://localhost:4434 identities list -f=json-pretty