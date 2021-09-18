#!/bin/bash
IDENTITIES=$(curl -s -X GET "http://localhost:4434/identities")
if [ $? -ne 0 ]; then
    echo "Filed to get identities from Kratos"
    exit 1
fi

IDENTITY=$(jq "map(select(.traits.email == \"$1\")) | .[0]" <<< "$IDENTITIES")
if [ "$IDENTITY" == "null" ]; then
    echo "Could not find identity with email $1"
    exit 1
fi

echo "Adding tenant $2 to identity with email $1"
UPDATED_IDENTITY=$(jq ".traits.tenants += [\"$2\"]" <<< "$IDENTITY")
if [ $? -ne 0 ]; then
    exit 1
fi

IDENTITY_ID=$(jq '.id' <<< "$IDENTITY")
KRATOS_UPDATE=$(jq '{ schema_id, traits }' <<< "$UPDATED_IDENTITY")

curl -f -s -X PUT "http://localhost:4434/identities/$IDENTITY_ID"       \
    -H 'Content-Type: application/json'                                 \
    -H 'Accept: application/json' --data @- <<< "$KRATOS_UPDATE" | jq
if [ $? -ne 0 ]; then
    exit 1
fi
