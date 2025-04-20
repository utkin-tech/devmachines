package main

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateDiskImage(baseImagePath, diskImagePath, diskSize string) error {
	if _, err := os.Stat(baseImagePath); os.IsNotExist(err) {
		return fmt.Errorf("base image %s does not exist", baseImagePath)
	}

	if _, err := os.Stat(diskImagePath); err == nil {
		fmt.Printf("disk image %s already exists, skipping creation\n", diskImagePath)
		return nil
	}

	cmd := exec.Command(
		"qemu-img",
		"create",
		"-b", baseImagePath,
		"-F", "qcow2",
		"-f", "qcow2",
		diskImagePath,
		diskSize,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create disk image: %v, output: %s", err, string(output))
	}

	if _, err := os.Stat(diskImagePath); os.IsNotExist(err) {
		return fmt.Errorf("command succeeded but disk image %s was not created", diskImagePath)
	}

	return nil
}
