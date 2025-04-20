package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func SetupNetworkBridge(bridgeName string, ethInterface string, tapInterface string, originalIP string) error {
	if err := runCommand("ip", "link", "add", "name", bridgeName, "type", "bridge"); err != nil {
		return fmt.Errorf("failed to create bridge: %v", err)
	}

	if err := runCommand("ip", "link", "set", "dev", bridgeName, "up"); err != nil {
		return cleanupBridge(bridgeName, fmt.Errorf("failed to bring bridge up: %v", err))
	}

	if err := runCommand("ip", "link", "set", ethInterface, "master", bridgeName); err != nil {
		return cleanupBridge(bridgeName, fmt.Errorf("failed to add %s to bridge: %v", ethInterface, err))
	}

	if err := runCommand("ip", "link", "set", "dev", ethInterface, "up"); err != nil {
		return cleanupBridge(bridgeName, fmt.Errorf("failed to bring %s up: %v", ethInterface, err))
	}

	if err := runCommand("ip", "tuntap", "add", "dev", tapInterface, "mode", "tap"); err != nil {
		return cleanupBridge(bridgeName, fmt.Errorf("failed to create tap interface: %v", err))
	}

	if err := runCommand("ip", "link", "set", tapInterface, "master", bridgeName); err != nil {
		return cleanupTap(tapInterface, cleanupBridge(bridgeName,
			fmt.Errorf("failed to add tap to bridge: %v", err)))
	}

	if err := runCommand("ip", "link", "set", tapInterface, "up"); err != nil {
		return cleanupTap(tapInterface, cleanupBridge(bridgeName,
			fmt.Errorf("failed to bring tap up: %v", err)))
	}

	if originalIP != "" {
		if err := runCommand("ip", "addr", "del", originalIP, "dev", ethInterface); err != nil {
			return cleanupTap(tapInterface, cleanupBridge(bridgeName,
				fmt.Errorf("failed to remove IP from %s: %v", ethInterface, err)))
		}
	}

	return nil
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s: %v, output: %s", strings.Join(cmd.Args, " "), err, string(output))
	}
	return nil
}

func cleanupBridge(bridgeName string, err error) error {
	_ = runCommand("ip", "link", "del", bridgeName)
	return err
}

func cleanupTap(tapName string, err error) error {
	_ = runCommand("ip", "link", "del", tapName)
	return err
}
