variable "ALPINE_VERSION" {
    default = null
}

group "default" {
    targets = ["alpine", "ubuntu"]
}

target "alpine" {
    dockerfile = "./alpine.Dockerfile"
    args = {
        ALPINE_VERSION = ALPINE_VERSION
    }
    tags = ["devmachines/kvmtool:alpine"]
}

target "ubuntu" {
    dockerfile = "./ubuntu.Dockerfile"
    tags = ["devmachines/kvmtool:ubuntu"]
}
