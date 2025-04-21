ARG ALPINE_VERSION="3.20"

FROM alpine:${ALPINE_VERSION}

WORKDIR /image

COPY images/ubuntu.img ubuntu.img
