# DevMachines

## Setup Diagram

```mermaid
flowchart TD
    User -->|User,Password,SSHKeys| SetupCloudInit
    User -->|VMDiskSize| SetupDisk
    User -->|CPU,RAM| StartVM
    SetupNetwork -->|Addresses,Gateway| SetupCloudInit
    SetupNetwork -->|NetworkParams| StartVM
    SetupDisk -->|DiskParams| StartVM
    SetupCloudInit -->|CloudInitParams| StartVM
```
