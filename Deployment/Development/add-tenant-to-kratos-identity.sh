#!/bin/bash
IDENTITY=$(docker exec dolittle_authentication_login-kratos_1 kratos --endpoint http://localhost:4434 identities get $1 -f=json)
if [ $? -ne 0 ]; then
    exit 1
fi

UPDATED_IDENTITY=$(jq ".traits.tenants += [\"$2\"]" <<< "$IDENTITY")
if [ $? -ne 0 ]; then
    exit 1
fi

KRATOS_UPDATE=$(jq '{ schema_id, traits }' <<< "$UPDATED_IDENTITY")

curl -f -s -X PUT "http://localhost:4434/identities/$1"                 \
    -H 'Content-Type: application/json'                                 \
    -H 'Accept: application/json' --data @- <<< "$KRATOS_UPDATE" | jq
if [ $? -ne 0 ]; then
    exit 1
fi