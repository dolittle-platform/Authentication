#!/bin/bash
curl -s -X GET "http://localhost:4434/admin/identities" | jq
