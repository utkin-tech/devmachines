# OverlayFS

## Start container
```sh
docker run --privileged -it --rm --mount type=image,source=ubuntu,target=/rootfs ubuntu
```

## Create OverlayFS

Create infrastucture for OverlayFS. It very important place `upper` and `work` folders inside `tmpfs`. Without this action nothing will works.
```sh
mkdir -p /tmp/overlay
mount -t tmpfs tmpfs /tmp/overlay
mkdir -p /tmp/overlay/upper
mkdir -p /tmp/overlay/work
mkdir -p /rootfs
mount -t overlay overlay -o lowerdir=/image,upperdir=/tmp/overlay/upper,workdir=/tmp/overlay/work /rootfs
```
