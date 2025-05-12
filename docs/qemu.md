# QEMU

## Start

Sample command:
```sh
qemu-system-x86_64  \
  -machine accel=kvm,type=q35 \
  -cpu host \
  -m 2G \
  -nographic \
  -device virtio-net-pci,netdev=net0 \
  -netdev user,id=net0,hostfwd=tcp::2222-:22 \
  -drive if=virtio,format=qcow2,file=disk.img \
  -drive if=virtio,format=raw,file=seed.iso \
  -vga std \
  -vnc 0.0.0.0:0
```

## Serial

Complex arguments:
```sh
-chardev socket,id=serial0,path=/tmp/serial0.sock,server=on,wait=off \
-device isa-serial,chardev=serial0 \
```

Simple argument (current):
```sh
-serial unix:/socks/serial.sock,server,nowait \
```

Connect to socket:
```sh
socat STDIO,cfmakeraw UNIX:/socks/serial.sock
```

## VNC

Unix socket arguments:
```sh
-vga std \
-vnc unix:/socks/vnc.sock \
```

## Minimal

```sh
qemu-system-x86_64 \
  -enable-kvm -smp 2 \
  -m 2G \
  -nographic \
  -drive if=virtio,file=disk2.img \
  -drive file=cloudinit.iso,format=raw,if=virtio
```
