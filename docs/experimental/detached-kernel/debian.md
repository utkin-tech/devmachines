Create a raw disk image:

```sh
qemu-img create -f qcow2 debian.qcow2 8G
```

Format the disk image with a filesystem:

```sh
sudo modprobe nbd max_part=8
sudo qemu-nbd --connect=/dev/nbd0 debian.qcow2
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
sudo mkdir -p /mnt/debian
sudo mount /dev/nbd0p1 /mnt/debian
```

Install utils:

```sh
sudo apt-get install debootstrap
```

Install a base Debian system using debootstrap:

```sh
sudo debootstrap --arch=amd64 bookworm /mnt/debian http://deb.debian.org/debian
```

Configure the basic system:

```sh
sudo chroot /mnt/debian

# Set root password
passwd

# Install necessary packages
apt-get update
apt-get install -y openssh-server sudo

# Configure networking
echo "debian-qemu" > /etc/hostname
echo "127.0.0.1 localhost" > /etc/hosts
echo "127.0.1.1 debian-qemu" >> /etc/hosts

# Create /etc/fstab
echo "/dev/sda1 / ext4 defaults 0 1" > /etc/fstab

# Exit chroot
exit
```

Unmount and disconnect:

```sh
sudo umount /mnt/debian
sudo qemu-nbd --disconnect /dev/nbd0
```

Start VM:

```sh
qemu-system-x86_64 \
  -enable-kvm -m 2G -smp 2 \
  -kernel /boot/vmlinuz-$(uname -r) \
  -initrd /boot/initrd.img-$(uname -r) \
  -append "root=/dev/sda1 rw console=ttyS0" \
  -drive file=debian.qcow2,format=qcow2 \
  -nic user,hostfwd=tcp::2222-:22 \
  -nographic
```

Get IP address:

```sh
sudo dhclient ens3
```
