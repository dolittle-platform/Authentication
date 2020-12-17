FROM golang:latest as build
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go build .

FROM scratch
COPY --from=build /build/impersonator /bin/impersonator
COPY .docker/group /etc/group
COPY .docker/passwd /etc/passwd
COPY .docker/nsswitch.conf /etc/nsswitch.conf
COPY --chown=impersonator:impersonator .docker/config.yaml /var/lib/impersonator/.impersonator.yaml
USER impersonator
ENTRYPOINT [ "/bin/impersonator" ]
CMD [ "serve" ]