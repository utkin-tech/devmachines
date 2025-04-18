group "default" {
    targets = ["runtime", "ubuntu"]
}

target "runtime" {
    dockerfile = "runtime.Dockerfile"
    tags = ["devmachines/runtime"]
}

target "ubuntu" {
    dockerfile = "ubuntu.Dockerfile"
    tags = ["devmachines/ubuntu"]
}
