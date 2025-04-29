variable "GO_VERSION" {
    default = null
}

variable "ALPINE_VERSION" {
    default = null
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
    context = "./image/ubuntu/24.04"
    tags = ["devmachines/ubuntu:24.04", "devmachines/ubuntu:noble"]
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


target "exit0" {
    context = "./image"
    tags = ["devmachines/exit0"]
}
