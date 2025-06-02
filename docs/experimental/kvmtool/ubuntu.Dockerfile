FROM ubuntu

RUN apt update && apt install -y build-essential git zlib1g-dev libaio-dev

RUN git clone git://git.kernel.org/pub/scm/linux/kernel/git/will/kvmtool.git

WORKDIR /kvmtool

RUN make
