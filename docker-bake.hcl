variable "GO_VERSION" {
    default = null
}

variable "ALPINE_VERSION" {
    default = null
}

variable "UBUNTU_VERSION" {
    default = 24.04
}

variable "UBUNTU_CODENAME" {
    default = "noble"
}

variable "DESTDIR" {
    default = "./bin"
}

target "_common" {
    args = {
        GO_VERSION = GO_VERSION
        ALPINE_VERSION = ALPINE_VERSION
        BUILDKIT_CONTEXT_KEEP_GIT_DIR = 1
    }
}

group "default" {
    targets = ["runtime", "ubuntu"]
}

target "runtime" {
    inherits = ["_common"]
    tags = ["devmachines/runtime"]
}

target "ubuntu" {
    context = "./images/image/ubuntu/${UBUNTU_VERSION}"
    tags = ["devmachines/ubuntu:${UBUNTU_VERSION}", "devmachines/ubuntu:${UBUNTU_CODENAME}"]
}

target "alpine" {
    context = "./images/rootfs/alpine"
    args = {
        ALPINE_VERSION = ALPINE_VERSION
    }
    tags = ["devmachines/alpine"]
}

target "binariy" {
    inherits = ["_common"]
    target = "binary"
    output = ["${DESTDIR}/build"]
}

target "novnc" {
    context = "./novnc"
    tags = ["devmachines/novnc"]
}
