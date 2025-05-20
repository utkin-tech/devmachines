# VirtioFS inside container

## Installation

Install virtiofsd
```sh
apk add virtiofsd
```

## Setup

Start virtiofsd
```sh
/usr/libexec/virtiofsd --socket-path=/tmp/vhostqemu --shared-dir /rootfs --cache auto --sandbox none
```

Start VM used vitiofs as rootfs:
```sh
qemu-system-x86_64 \
  -enable-kvm \
  -smp 2 \
  -m 2G \
  -object memory-backend-memfd,id=mem,size=2G,share=on \
  -numa node,memdev=mem \
  -kernel /kernel/bzImage2 \
  -chardev socket,id=char0,path=/tmp/vhostqemu \
  -device vhost-user-fs-pci,queue-size=1024,chardev=char0,tag=myfs \
  -append "rootfstype=virtiofs root=myfs rw console=ttyS0" \
  -nic user,hostfwd=tcp::2222-:22 \
  -nographic
```

Add cloud-init option:
```sh
  -drive if=virtio,format=raw,file=/disks/cloudinit.iso
```

## Links

- https://virtio-fs.gitlab.io/howto-boot.html
- https://virtio-fs.gitlab.io/howto-qemu.html
- https://gitlab.com/virtio-fs/virtiofsd
