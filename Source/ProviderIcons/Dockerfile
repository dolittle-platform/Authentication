FROM nginx:1.21.6
ARG BASE_PATH="/.auth/provider-icons"

RUN rm -r /usr/share/nginx/html
ENV ROOT_PATH="/usr/share/nginx/html${BASE_PATH}"

COPY ./github.svg ${ROOT_PATH}/github.svg
COPY ./microsoft.svg ${ROOT_PATH}/microsoft.svg
