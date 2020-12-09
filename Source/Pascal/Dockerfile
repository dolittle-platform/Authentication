FROM golang:latest as build
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go build .

FROM alpine:latest as ssl-certs

FROM scratch
COPY --from=ssl-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /build/pascal /bin/pascal
COPY .docker/group /etc/group
COPY .docker/passwd /etc/passwd
COPY .docker/nsswitch.conf /etc/nsswitch.conf
COPY --chown=pascal:pascal .docker/config.yaml /var/lib/pascal/.pascal.yaml
USER pascal
ENTRYPOINT [ "/bin/pascal" ]
CMD [ "serve" ]