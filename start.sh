#!/bin/bash

BASE_IMAGE="/image/ubuntu.img"
DISK_IMAGE="/blobs/disk.img"
DISK_SIZE="10G"

if [ ! -f "$BASE_IMAGE" ]; then
    echo "Error: Base image $BASE_IMAGE does not exist!"
    exit 1
fi

if [ -f "$DISK_IMAGE" ]; then
    echo "Disk image $DISK_IMAGE already exists. Skipping creation."
else
    echo "Creating new disk image from base image..."
    qemu-img create -b "$BASE_IMAGE" -F qcow2 -f qcow2 "$DISK_IMAGE" "$DISK_SIZE"
    
    if [ $? -eq 0 ] && [ -f "$DISK_IMAGE" ]; then
        echo "Successfully created disk image $DISK_IMAGE"
    else
        echo "Error: Failed to create disk image!"
        exit 1
    fi
fi

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
