ARG ALPINE_VERSION="3.20"

FROM alpine:${ALPINE_VERSION}

RUN apk add alpine-base cloud-init

RUN --mount=type=bind,source=setup.sh,target=setup.sh \
    sh setup.sh

RUN echo "root:pass" | chpasswd
