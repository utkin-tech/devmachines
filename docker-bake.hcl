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
    dockerfile = "runtime.Dockerfile"
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
    context = "./image"
    tags = ["devmachines/ubuntu"]
}

target "binariy" {
    inherits = ["_common"]
    target = "binary"
    output = ["${DESTDIR}/build"]
}
