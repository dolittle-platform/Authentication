FROM golang:latest as build-backend
WORKDIR /build
COPY ./Backend /build
ENV CGO_ENABLED=0
RUN go build .

FROM node:latest as build-web
WORKDIR /build
COPY ./Web/package.json /build/package.json
RUN yarn
COPY ./Web /build
RUN yarn build

FROM scratch
COPY --from=build-backend /build/login /bin/login
COPY .docker/group /etc/group
COPY .docker/passwd /etc/passwd
COPY --from=build-web /build/wwwroot /var/lib/login/wwwroot
COPY --chown=login:login .docker/config.yaml /var/lib/login/.login.yaml
USER login
WORKDIR /var/lib/login
ENTRYPOINT [ "/bin/login" ]
CMD [ "serve" ]