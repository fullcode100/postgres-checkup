FROM alpine:3.9

RUN apk add --update --no-cache \
    bash \
    openssh-client \
    postgresql-client \
    coreutils \
    jq \
    go \
    gawk \
    sed \
    make \
    build-base

COPY . .
