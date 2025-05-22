ARG GO_VERSION="1.24.2"
ARG ALPINE_VERSION="3.20"

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
ENV GO111MODULE=auto
ENV CGO_ENABLED=0
WORKDIR /src

RUN --mount=type=bind,src=go.mod,target=go.mod \
    --mount=type=bind,src=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG LDFLAGS="-s -w"
ARG BUILDTAGS=""
RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache \
    go build -tags "${BUILDTAGS}" -trimpath -ldflags "${LDFLAGS}" -o "/devmachines-runtime" ./cmd/devmachines-runtime

FROM scratch AS binary
COPY --from=build /devmachines-runtime /devmachines-runtime

FROM alpine:${ALPINE_VERSION}

RUN apk add qemu-system-x86_64 qemu-img iproute2 cdrkit

COPY --link --from=binary /devmachines-runtime /devmachines-runtime

COPY ./static/ /static/

COPY --link --from=devmachines/novnc /vnc_lite.html /static/vnc/index.html

CMD ["/devmachines-runtime"]

EXPOSE 22/tcp 8080/tcp 8081/tcp

VOLUME /socks /disks
