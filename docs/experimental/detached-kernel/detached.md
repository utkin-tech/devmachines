# Detached kernel

Download ubuntu image:
```sh
wget https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-disk-kvm.img
cp jammy-server-cloudimg-amd64-disk-kvm.img disk.img
```

Run QEMU with detached kernel:
```sh
qemu-system-x86_64 \
  -enable-kvm -smp 2 \
  -m 2G \
  -nographic \
  -enable-kvm \
  -drive if=virtio,file=disk.img \
  -kernel bzImage \
  -append "root=/dev/vda1 console=ttyS0"
```
