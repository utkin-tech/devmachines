ARG ALPINE_VERSION="3.20"

FROM alpine:${ALPINE_VERSION}

RUN apk add build-base git make linux-headers libaio-dev zlib-dev

RUN git clone git://git.kernel.org/pub/scm/linux/kernel/git/will/kvmtool.git

WORKDIR /kvmtool

RUN sed -i '1i\#include <libgen.h>' vfio/core.c

RUN make
