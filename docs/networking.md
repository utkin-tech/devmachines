# Networking

This section describes the networking models and practices used in `devmachines/runtime`.

## Creating a TAP Interface Inside a Container

This example demonstrates how to create a TAP interface inside a Docker container.

### Start the Container
```sh
docker run -it \
  --name iptest \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  --device /dev/kvm \
  alpine
```

### Install Required Tools
```sh
apk add iproute2
```

### Configure Bridge and TAP Interfaces
```sh
# Create bridge and bring it up
ip link add name br0 type bridge
ip link set dev br0 up

# Add existing interface to bridge
ip link set eth0 master br0

# Create and configure TAP interface
ip tuntap add dev tap0 mode tap
ip link set tap0 master br0
ip link set dev tap0 up

# Clean up existing configuration
ip addr flush dev eth0
```

### Debugging Configuration (Optional)
```sh
# Assign IP address to bridge
ip addr add 172.17.0.2/16 dev br0
ip route add default via 172.17.0.1
```
