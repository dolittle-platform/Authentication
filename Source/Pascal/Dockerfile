FROM golang:latest as build
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go build .

FROM scratch
COPY --from=build /build/pascal /bin/pascal
COPY .docker/group /etc/group
COPY .docker/passwd /etc/passwd
COPY --chown=pascal:pascal .docker/config.yaml /var/lib/pascal/.pascal.yaml
USER pascal
ENTRYPOINT [ "/bin/pascal" ]
CMD [ "serve" ]