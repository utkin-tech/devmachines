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


## Inspect Docker image content

```sh
docker run -ti --rm  -v /var/run/docker.sock:/var/run/docker.sock docker.io/wagoodman/dive devmachines/ubuntu
```
