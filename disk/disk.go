package disk

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	BaseImagePath = "/image/ubuntu.img"
	DiskImagePath = "/blobs/disk.img"
)

type Storage interface {
	Size() string
}

func SetupDisk(storage Storage) ([]string, error) {
	err := createDiskImage(BaseImagePath, DiskImagePath, storage.Size())
	if err != nil {
		return nil, err
	}

	return []string{
		"-drive", fmt.Sprintf("file=%s,format=qcow2,if=virtio", DiskImagePath),
	}, nil
}

func createDiskImage(baseImagePath string, diskImagePath string, diskSize string) error {
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

	return nil
}
