docker run --name qemu -it --privileged --mount type=bind,source="$(pwd)"/blobs,target=/blobs alpine

apk add qemu-system-x86_64

qemu-system-x86_64 \
  -m 2048 \
  -smp 2 \
  -enable-kvm \
  -drive file=/blobs/vm-disk.img,format=qcow2,if=virtio \
  -drive file=/blobs/seed.iso,format=raw,if=virtio \
  -netdev user,id=net0,hostfwd=tcp::2222-:22 \
  -device virtio-net-pci,netdev=net0 \
  -nographic
