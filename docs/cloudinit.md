# Cloud Init

In `devmachimes/runtime`, [Cloud-init](https://cloudinit.readthedocs.io/en/latest/#) is used to pass setup parameters when starting a VM. The process of creating cloud-init is handled by the `SetupCloudInit` function, as described in the `Setup Diagram`.

In the current version, `SetupCloudInit` generates a `/disks/cloudinit.iso` file using the `genisoimage` command. As specified in the cloud-init documentation, this ISO contains the following files: `meta-data`, `user-data`, and `network-config`.

During the `SetupCloudInit` process, these files are first stored in the OS temporary directory (`/tmp` by default). Then, `genisoimage` is executed. Here’s an example of the command:

```sh
genisoimage \
    -o /disks/cloudinit.iso \
    -volid cidata \
    -joliet -rock \
    /tmp/meta-data \
    /tmp/user-data \
    /tmp/network-config
```

## Meta Data

The purpose of this file is to instruct cloud-init to initialize the VM.

Example content of the `meta-data` file:
```
instance-id: b1b5cfbbff6a707cd9d83d4386b8f107921b8ff0
local-hostname: my-vm
```

The runtime changes the `instance-id` every time the container starts. This is a workaround to ensure configuration changes in `devmachines/runtime` are applied inside the VM.

## User Data

This file provides information about the default user to be created in the VM, including parameters such as username, password, SSH keys, VM hostname, and others.

Example content of the `user-data` file:
```
#cloud-config
hostname: my-vm
manage_etc_hosts: true
fqdn: my-vm
user: user
password: pass
ssh_authorized_keys:
    - ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICeyDWnYtatyEiHKjT/jpLUA7FLMoe97ItrkaLA2cYsX qemu_machine
chpasswd:
    expire: false
users:
    - default
```

Currently, an unencrypted password (`pass`) is provided in the config, so we highly recommend not using this in production. Future updates plan to offer an option to set up the VM without a password.

## Network Data

This file provides network configuration parameters to the VM. `devmachines/runtime` retrieves network information from the container's network interface. For more details, see the `SetupNetwork` section.

Example content of the `network-config` file:
```
version: 2
ethernets:
    ens3:
        dhcp4: false
        addresses:
            - 172.18.0.2/16
        gateway4: 172.18.0.1
        nameservers:
            addresses:
                - 8.8.8.8
                - 1.1.1.1
```

By default, `8.8.8.8` and `1.1.1.1` are set as DNS servers, and `ens3` is configured as the network interface inside the VM.

## Exploring the Content

You can inspect the contents of `cloudinit.iso` using the following commands.

To mount the disk to `/mnt/iso`:
```sh
sudo mkdir /mnt/iso
sudo mount -o loop disks/seed.iso /mnt/iso
```

To unmount `cloudinit.iso`:
```sh
sudo umount /mnt/iso
```
