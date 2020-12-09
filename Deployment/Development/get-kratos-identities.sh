#!/bin/bash
docker exec dolittle_authentication_login-kratos_1 kratos           \
    --endpoint http://localhost:4434 identities list -f=json-pretty