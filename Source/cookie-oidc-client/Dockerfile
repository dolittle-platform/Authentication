FROM golang:latest as build
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go build .

FROM scratch
COPY --from=build /build/cookie-oidc-client /bin/cookie-oidc-client
COPY .docker/group /etc/group
COPY .docker/passwd /etc/passwd
COPY --chown=cookie-oidc-client:cookie-oidc-client .docker/config.yaml /var/lib/cookie-oidc-client/.cookie-oidc-client.yaml
USER cookie-oidc-client
ENTRYPOINT [ "/bin/cookie-oidc-client" ]
CMD [ "serve" ]