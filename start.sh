# TODO: Add dynamic disk generation if not exists
qemu-img create -b /image/ubuntu.img -F qcow2 -f qcow2 /blobs/disk.img 10G

ip link add name br0 type bridge
ip link set dev br0 up

ip link set eth0 master br0
ip link set dev eth0 up

ip tuntap add dev tap0 mode tap
ip link set tap0 master br0
ip link set tap0 up

# TODO: Add dynamic adress resolution
ip addr del 172.20.0.2/16 dev eth0

# TODO: Add params selection from env variables
# TODO: Add dynamic seed.iso generation
qemu-system-x86_64 \
  -m 2048 \
  -smp 2 \
  -enable-kvm \
  -drive file=/blobs/disk.img,format=qcow2,if=virtio \
  -drive file=/blobs/seed.iso,format=raw,if=virtio \
  -netdev tap,id=net0,ifname=tap0,script=no,downscript=no \
  -device virtio-net-pci,netdev=net0 \
  -nographic
