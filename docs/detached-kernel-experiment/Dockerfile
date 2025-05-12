ARG ALPINE_VERSION="3.20"

FROM alpine:${ALPINE_VERSION}

RUN apk add qemu-system-x86_64 qemu-img iproute2 cdrkit virtiofsd

COPY --link --from=binary /devmachines-runtime /devmachines-runtime

CMD ["/devmachines-runtime"]
