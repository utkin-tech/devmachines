# RootFS

## Creation

Install debootstrap:
```sh
sudo apt update
sudo apt install debootstrap
```

Create a Directory for the Rootfs:
```sh
export ROOTFS_DIR=ubuntu_rootfs
mkdir -p $ROOTFS_DIR
```

Run debootstrap to Install a Minimal Ubuntu:
```sh
sudo debootstrap --arch=amd64 jammy $ROOTFS_DIR http://archive.ubuntu.com/ubuntu/
```

## Add user

To set a password for the **root** user in `/etc/shadow`, follow these steps carefully. The `/etc/shadow` file stores encrypted password hashes, and you should never edit it directly with a text editor.

### Generate a Password Hash

Generate password for user:
```sh
openssl passwd -6 -salt $(openssl rand -hex 4) "yourpassword"
```

This will output a hashed string like:
```
$6$salt$hashedpassword
```

### Edit `/etc/shadow`
- Open `/etc/shadow` with `sudo vipw -s` (safer than a plain editor) or:
  ```bash
  sudo nano /etc/shadow
  ```
- Find the root entry (first line):
  ```
  root:!:19164:0:99999:7:::
  ```
- Replace `!` or `*` with the generated hash:
  ```
  root:$6$salt$hashedpassword:19164:0:99999:7:::
  ```
- **Save and exit** (in `nano`: `Ctrl+O`, `Enter`, `Ctrl+X`).
