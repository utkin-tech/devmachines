ARG ALPINE_VERSION="3.20"

FROM alpine:${ALPINE_VERSION} AS build

RUN apk add gcc

WORKDIR /src

COPY exit0.c exit0.c

RUN gcc -static -nostdlib -fno-builtin -o /exit0 exit0.c

FROM scratch

COPY --link --from=build /exit0 /exit0

CMD ["/exit0"]
