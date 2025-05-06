# Kernel

## Compilation

Install utils:
```sh
sudo apt-get install git fakeroot build-essential ncurses-dev xz-utils libssl-dev bc flex libelf-dev bison
```

Get sources:
```sh
git clone https://github.com/torvalds/linux.git
cd linux/
```

Setup default config:
```sh
make x86_64_defconfig
make kvm_guest.config
```

Enter menuconfig:
```sh
make menuconfig
```

Enable FUSE and Virtio Filesystem (both):
```
 -> File systems
   -> FUSE (Filesystem in Userspace) support [*]
   -> Virtio Filesystem (VIRTIO_FS) [*]
```

Compile kernel:
```sh
make -j8
# OR
make -j`nproc`
```

Copy kernel in working directory
```sh
cp arch/x86/boot/bzImage /home/user
```

## Links

- https://phoenixnap.com/kb/build-linux-kernel
- https://vccolombo.github.io/cybersecurity/linux-kernel-qemu-setup/
- https://gist.github.com/ncmiller/d61348b27cb17debd2a6c20966409e86
