#!/bin/bash

openssl req -x509            \
    -nodes                   \
    -days 730                \
    -newkey rsa:2048         \
    -keyout certificate.key  \
    -out certificate.crt     \
    -config certificate.conf \
    -extensions 'v3_req'
