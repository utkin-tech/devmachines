# kvmtool

## Build

Add build utils (alpine):
```sh
apk add build-base git make linux-headers libaio-dev zlib-dev
```

Fix for `alpine`. If there is error when build add line `#include <libgen.h>` into `vfio/core.c` file.
```sh
sed -i '1i\#include <libgen.h>' vfio/core.c
```

Clone repo:
```sh
git clone git://git.kernel.org/pub/scm/linux/kernel/git/will/kvmtool.git
cd kvmtool
```
:
Compile:
```sh
make
```
