variable "GO_VERSION" {
    default = null
}

variable "DESTDIR" {
    default = "./bin"
}

variable "IMAGE" {
    default = "erlnby/yadro-app"
}

target "_common" {
    dockerfile = "runtime.Dockerfile"
    args = {
        GO_VERSION = GO_VERSION
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
    dockerfile = "ubuntu.Dockerfile"
    tags = ["devmachines/ubuntu"]
}

target "binariy" {
    inherits = ["_common"]
    target = "binary"
    output = ["${DESTDIR}/build"]
}
