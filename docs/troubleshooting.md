# Troubleshooting

## Start QEMU inside container

Test KVM:
```sh
qemu-system-x86_64 \
    -m 2048 \
    -smp 2 \
    -enable-kvm \
    -nographics
```

Start VM:
```sh
qemu-system-x86_64 \
    -m 2048 \
    -smp 2 \
    -drive file=/disks/disk.img,format=qcow2,if=virtio \
    -drive file=/disks/cloudinit.iso,format=raw,if=virtio \
    -netdev tap,id=net0,ifname=tap0,script=no,downscript=no \
    -device virtio-net-pci,netdev=net0 \
    -qmp unix:/tmp/qmp-sock,server,wait=off \
    -enable-kvm \
    -nographic
```

## Connect to serial terminal

Install `socat`:
```sh
sudo apt install socat
```

```sh
socat -,raw,echo=0,escape=0x18 UNIX-CONNECT:./socks/serial.sock
```

Exit keys `Ctrl + X`

## Inspect Docker image content

```sh
docker run -ti --rm  -v /var/run/docker.sock:/var/run/docker.sock docker.io/wagoodman/dive devmachines/ubuntu
```

## Inspect Qcow2 image

```sh
sudo modprobe nbd
```

```sh
sudo qemu-nbd --connect=/dev/nbd0 file.qcow2
```

```sh
sudo fdisk -l /dev/nbd0
```

```sh
sudo mount /dev/nbd0p1 /mnt
```

```sh
sudo umount /mnt
sudo qemu-nbd --disconnect /dev/nbd0
```

## Inspect ISO

Make `raw` from `qcow2` (optional):
```sh
qemu-img convert -f qcow2 -O raw file.qcow2 file.raw
```

```sh
sudo mount -o loop file.raw /mnt
```

```sh
sudo umount /mnt
```
