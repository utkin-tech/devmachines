package main

import (
	"fmt"
	"io"
	"os/exec"
)

// QEMUConfig holds configuration parameters for the QEMU virtual machine
type QEMUConfig struct {
	MemoryMB int       // Memory in megabytes (e.g., 2048)
	CPUCores int       // Number of CPU cores (e.g., 2)
	Output   io.Writer // Where to direct output (nil for default)
	Wait     bool      // Whether to wait for VM to exit
}

func StartVM(config QEMUConfig, extraArgs ...string) error {
	args := []string{
		"-m", fmt.Sprintf("%d", config.MemoryMB),
		"-smp", fmt.Sprintf("%d", config.CPUCores),
		"-enable-kvm",
		"-nographic",
	}

	args = append(args, extraArgs...)

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
