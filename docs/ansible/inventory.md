# Ansible Dynamic Inventory Plugin for DevMachines

An Ansible dynamic inventory plugin for DevMachines, heavily based on the `community.docker.docker_containers` inventory plugin.

## Ansible Installation

Follow these steps to install Ansible using `pipx`. For more information, refer to the official [Ansible documentation](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx).

Install `pipx` (Ubuntu):
```sh
sudo apt install pipx
```

Install Ansible with `pipx`:
```sh
pipx install --include-deps ansible
```

## Usage

Install the `devmachines.core` Ansible collection:
```sh
ansible-galaxy collection install devmachines.core
```

Create a `devmachines.yml` file with the following content:
```yaml
---
plugin: devmachines.core.devmachines
```

Get VMs list:
```sh
ansible-inventory --list -i devmachines.yml
```

**Note:** For proper functionality, containers must meet these requirements:
- Have the label `tech.utkin.devmachines.runtime=true`
- Have `22/tcp` in `NetworkSettings.Ports` connected to the VM's SSH port

## Development

Install the `community.docker` plugin:
```sh
ansible-galaxy collection install community.docker
```

Add the `requests` package (required by the `community.docker.docker_containers` plugin):
```sh
pipx inject ansible requests
```

Create an `ansible.cfg` file with the following content:
```ini
[defaults]
inventory_plugins = ./ansible/plugins/inventory
```

Create a `devmachines.yml` file with the following content:
```yaml
# devmachines.yml
plugin: devmachines
```

## Build

Build collection:
```sh
ansible-galaxy collection build --force
```

Install collection:
```sh
ansible-galaxy collection install --force devmachines-core-0.1.0.tar.gz
```

## References

- [Docker Containers Inventory Plugin Documentation](https://docs.ansible.com/ansible/latest/collections/community/docker/docker_containers_inventory.html)
- [community.docker.docker_containers Source Code](https://github.com/ansible-collections/community.docker/blob/main/plugins/inventory/docker_containers.py)
