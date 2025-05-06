# DevMachines: QEMU Virtual Machines in Docker Containers

**Run full-featured virtual machines (VMs) inside Docker containers**—combining the isolation of virtualization with the portability of containers. Ideal for development, testing, CI/CD, embedded systems, and security research.

## 🔥 Key Features
✅ **VM-as-a-Container** – Manage VMs like containers (`docker run`, `docker compose`)
✅ **Configurable Resources** – Set CPU cores, RAM, and disk size via environment variables
✅ **Snapshot-Based Storage** – Ephemeral or persistent disk options
✅ **SSH & User Management** – Pre-configured users, passwords, and SSH keys
✅ **Cloud-Init Support** – Customize VM boot behavior (networking, packages, users)
✅ **Networking Flexibility** – Bridge, TAP, or NAT networking for advanced use cases

## 🚀 Quick Start
1. Download the `compose.yaml`:
   ```sh
   wget https://raw.githubusercontent.com/utkin-tech/devmachines/main/compose.yaml
   docker compose up -d
   ```
2. Connect via SSH:
   ```sh
   VM_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{break}}{{end}}' runtime)
   ssh user@"$VM_IP"
   ```

## 📌 Use Cases
- **Developers**: Disposable dev environments with kernel/module testing
- **CI/CD**: Isolated build/test runners with VM-level reproducibility
- **Security Research**: Sandbox untrusted code in lightweight VMs
- **Embedded/IoT**: Emulate ARM/x86 devices in containers

## ⚙️ Configuration
| Env        | Example          | Validation             | Default | Description                      |
| ---------- | ---------------- | ---------------------- | ------- | -------------------------------- |
| `CPU`      | `2`              | required               |         | CPU cores                        |
| `MEMORY`   | `4G`             | required, validBytes   |         | RAM                              |
| `STORAGE`  | `20G`            | required, validBytes   |         | Disk size                        |
| `USER`     | `dev`            | required               |         | Default username                 |
| `PASSWORD` | `secret`         | required               |         | Default password                 |
| `SSH_KEYS` | `ssh-ed25519...` | required               |         | Comma-separated public SSH keys  |
| `NETWORK`  |                  | required, validNetwork | `NAT`   | Network type                     |
| `VNC`      | `:77`            |                        |         | QEMU-style argument for `-vnc`   |
| `PORTS`    | `2222:22`        | validPorts             |         | Read [Valid Ports](#valid-ports) |

### Valid Bytes

The values for `MEMORY` and `STORAGE` must be human-readable strings representing sizes in bytes. Supported suffixes include: `K` (kibibytes), `M` (mebibytes), `G` (gibibytes), and `T` (tebibytes). Units are case-insensitive and the `b` suffix is optional. For validation used `units.RAMInBytes` from `github.com/docker/go-units`.

### Valid Network

Valid values for `NETWORK` are `NAT` and `BRIDGE`.

### Valid Ports

Comma-separated Docker-style ports mapping. Have effect only when `NETWORK` is `NAT`. For validation used `nat.ParsePortSpec` from `github.com/docker/go-connections/nat`.

## 📂 Docs & Architecture

### System Diagram

```mermaid
%%{init: {"flowchart": {"defaultRenderer": "elk"}} }%%
flowchart TD
    RC[[RuntimeContainer]] -->|NetworkParams| GetNetworkInfo
    RC --> RCE([RuntimeContainerEnv])

    GetNetworkInfo --> SetupNetwork
    GetNetworkInfo --> SetupCloudInit

    RCE -->|USER,PASSWORD,SSH_KEYS| SetupCloudInit
    RCE -->|STORAGE| SetupDisk
    RCE -->|CPU,MEMORY| StartVM
    RCE -.->|VNC| SetupVNC
    RCE -->|NETWORK| GetNetworkInfo

    subgraph Args Providers
    SetupNetwork
    SetupCloudInit
    SetupVNC
    SetupDisk
    end

    %% Start VM Args
    SetupNetwork -->|Network Args| StartVM
    SetupDisk -->|Disk Args| StartVM
    SetupCloudInit -->|CloudInit Args| StartVM
    SetupVNC -.->|VNC Args| StartVM

    %% Disks
    Disk[(Disk)] --> SetupDisk
    BaseDisk[(Base Disk)] --> SetupDisk
    BaseDisk ---> Disk

```

### Details
- [Networking Setup](/docs/networking.md) (Bridge/TAP)
- [Cloud-Init Integration](/docs/cloudinit.md)

## ❓ Why?
- **No Hypervisor Overhead**: Uses Docker’s native capabilities + KVM acceleration
- **Declarative Management**: Define VMs in `docker-compose.yml` like containers
- **Cross-Platform**: Works anywhere Docker runs (Linux, macOS/WSL2, cloud)

Yes! Adding licensing, author information, and contribution guidelines makes your project more transparent and professional. Here’s how to integrate these details into your repo description (typically in the `README.md`):

## 📜 License
**DevMachines** is released under the **[MIT License](/LICENSE)**.

## 👨‍💻 Author
**Daniil Utkin** ([@erlnby](https://github.com/erlnby))
✉️ Email: [zolotoie@gmail.com](mailto:zolotoie@gmail.com)

## 🤝 Contributing
We welcome contributions! See [CONTRIBUTING.md](/CONTRIBUTING.md) for detailed guidelines.
