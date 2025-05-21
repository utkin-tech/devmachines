# Ansible Dynamic Inventory Plugin for DevMachines

An Ansible dynamic inventory plugin for DevMachines, heavily based on the `community.docker.docker_containers` inventory plugin.

## Ansible Installation

Follow these steps to install Ansible using `pipx`. For more information, refer to the official [Ansible documentation](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx).

1. Install `pipx` (Ubuntu):
   ```sh
   sudo apt install pipx
   ```

2. Install Ansible with `pipx`:
   ```sh
   pipx install --include-deps ansible
   ```

3. Install the `community.docker` plugin:
   ```sh
   ansible-galaxy collection install community.docker
   ```

4. Add the `requests` package (required by the `community.docker.docker_containers` plugin):
   ```sh
   pipx inject ansible requests
   ```

## Development

1. Create an `ansible.cfg` file with the following content:
   ```ini
   [defaults]
   inventory_plugins = ./plugins/inventory
   ```

2. Create a `devmachines.yml` file with the following content:
   ```yaml
   # devmachines.yml
   plugin: devmachines
   filters:
     - include: >-
         "tech.utkin.devmachines.runtime" in docker_config.Labels
     - exclude: true
   ```

**Note:** For proper functionality, containers must meet these requirements:
- Have the label `tech.utkin.devmachines.runtime=true`
- Have `22/tcp` in `NetworkSettings.Ports` connected to the VM's SSH port

## References

- [Docker Containers Inventory Plugin Documentation](https://docs.ansible.com/ansible/latest/collections/community/docker/docker_containers_inventory.html)
- [community.docker.docker_containers Source Code](https://github.com/ansible-collections/community.docker/blob/main/plugins/inventory/docker_containers.py)
