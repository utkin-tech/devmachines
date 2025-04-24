# DevMachines

## Setup Diagram

```mermaid
flowchart TD
    RCE -->|User,Password,SSHKeys| SetupCloudInit
    RCE -->|DiskSize| SetupDisk
    RCE -->|CPU,RAM| StartVM
    SetupNetwork -->|Addresses,Gateway| SetupCloudInit
    SetupNetwork -->|NetworkParams| StartVM
    SetupDisk -->|DiskParams| StartVM
    SetupCloudInit -->|CloudInitParams| StartVM
    RC[[RuntimeContainer]] -->|NetworkParams| SetupNetwork
    IC[[ImageContainer]] -->|BaseImage| SetupDisk
    RC --> RCE([RuntimeContainerEnv])
```

## Get container IP

```sh
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' runtime
```

## Create TAP interface inside container

Start container
```sh
docker run -it --name iptest --cap-add NET_ADMIN --device /dev/net/tun --device /dev/kvm alpine
```

Install iproute tool
```sh
apk add iproute2
```

Create bridge and tap interfaces
```sh
ip link add name br0 type bridge
ip link set dev br0 up
ip link set eth0 master br0
ip tuntap add dev tap0 mode tap
ip link set tap0 master br0
ip link set dev tap0 up
ip addr flush dev eth0
```

For debug add ip to bridge
```sh
ip addr add 172.17.0.2/16 dev br0
ip route add default via 172.17.0.1
```

## Start QEMU inside container

Test KVM
```sh
qemu-system-x86_64 \
    -m 2048 \
    -smp 2 \
    -enable-kvm \
    -nographics
```

Start VM
```sh
qemu-system-x86_64 \
    -m 2048 \
    -smp 2 \
    -drive file=/blobs/disk.img,format=qcow2,if=virtio \
    -drive file=/blobs/seed.iso,format=raw,if=virtio \
    -netdev tap,id=net0,ifname=tap0,script=no,downscript=no \
    -device virtio-net-pci,netdev=net0 \
    -qmp unix:/tmp/qmp-sock,server,wait=off \
    -enable-kvm \
    -nographic
```

## Inspect image content

```sh
docker run -ti --rm  -v /var/run/docker.sock:/var/run/docker.sock docker.io/wagoodman/dive devmachines/ubuntu
```
