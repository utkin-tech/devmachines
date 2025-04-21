package main

import (
	"fmt"
	"io"
	"os/exec"
)

func LaunchQEMUVM(config QEMUConfig) error {
	args := []string{
		"-m", fmt.Sprintf("%d", config.MemoryMB),
		"-smp", fmt.Sprintf("%d", config.CPUCores),
		"-enable-kvm",
		"-nographic",
	}

	cmd := exec.Command("qemu-system-x86_64", args...)
	cmd.Stdout = config.Output
	cmd.Stderr = config.Output

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start QEMU: %w", err)
	}

	if config.Wait {
		return cmd.Wait()
	}
	return nil
}

// QEMUConfig holds configuration parameters for the QEMU virtual machine
type QEMUConfig struct {
	MemoryMB      int       // Memory in megabytes (e.g., 2048)
	CPUCores      int       // Number of CPU cores (e.g., 2)
	SeedImagePath string    // Path to seed ISO (e.g., "/blobs/seed.iso")
	TapInterface  string    // Tap interface name (e.g., "tap0")
	Output        io.Writer // Where to direct output (nil for default)
	Wait          bool      // Whether to wait for VM to exit
}
