Create a raw disk image:

```sh
qemu-img create -f qcow2 alpine.qcow2 2G
```

Format the disk image with a filesystem:

```sh
sudo modprobe nbd max_part=8
sudo qemu-nbd --connect=/dev/nbd0 alpine.qcow2
sudo fdisk /dev/nbd0
```

In fdisk, create a new partition:

Press n for new partition
Press p for primary partition
Press 1 for partition number 1
Accept defaults for first and last sectors
Press a to make it bootable
Press w to write changes

Format and mount the partition:

```sh
sudo mkfs.ext4 /dev/nbd0p1
sudo mkdir -p /mnt/alpine
sudo mount /dev/nbd0p1 /mnt/alpine
```

Download Alpine minirootfs:

```sh
wget https://dl-cdn.alpinelinux.org/alpine/v3.21/releases/x86_64/alpine-minirootfs-3.21.3-x86_64.tar.gz
```

Unpack data in partition:

```sh
sudo tar -xzf alpine-minirootfs-3.21.3-x86_64.tar.gz -C /mnt/alpine
```

Configure the basic system:

```sh
sudo chroot /mnt/alpine /bin/sh

cat > /etc/resolv.conf <<EOF
nameserver 8.8.8.8
nameserver 8.8.4.4
EOF

# Update package indexes
apk update

# Install the basic packages
apk add openssh-server sudo

# Create /etc/fstab
echo "/dev/sda1 / ext4 defaults 0 1" > /etc/fstab

# Set up the network
echo "alpine-qemu" > /etc/hostname
echo "127.0.0.1 localhost" > /etc/hosts
echo "127.0.1.1 alpine-qemu" >> /etc/hosts

# Set up a root password
passwd root

# Enable SSH service when downloading
rc-update add sshd default

# Exit from chroot
exit
```

Unmount and disconnect:

```sh
sudo umount /mnt/alpine
sudo qemu-nbd --disconnect /dev/nbd0
```

Start VM:

```sh
qemu-system-x86_64 \
  -enable-kvm -m 2G -smp 2 \
  -kernel /boot/vmlinuz-$(uname -r) \
  -initrd /boot/initrd.img-$(uname -r) \
  -append "root=/dev/sda1 rw console=ttyS0" \
  -drive file=alpine.qcow2,format=qcow2 \
  -nic user,hostfwd=tcp::2222-:22 \
  -nographic
```
