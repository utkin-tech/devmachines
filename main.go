package main

import (
	"fmt"
	"os"

	"github.com/utkin-tech/devmachines/cloudinit"
	"github.com/utkin-tech/devmachines/disk"
	"github.com/utkin-tech/devmachines/network"
)

const InterfaceName = "eth0"

func main() {
	user := NewUser()

	net, err := network.NewNetwork(InterfaceName)
	if err != nil {
		fmt.Printf("failed to get info about network: %v\n", err)
		os.Exit(1)
	}

	diskArgs, err := disk.SetupDisk(user)
	if err != nil {
		fmt.Printf("failed to setup disk: %v\n", err)
		os.Exit(1)
	}

	cloudInitArgs, err := cloudinit.SetupCloudInit(net, user)
	if err != nil {
		fmt.Printf("failed to setup cloud-init: %v\n", err)
		os.Exit(1)
	}

	bridgeArgs, err := network.SetupBridge(net)
	if err != nil {
		fmt.Printf("failed to setup network bridge: %v\n", err)
		return
	}

	config := QEMUConfig{
		MemoryMB: 2048,
		CPUCores: 2,
		Output:   os.Stdout,
		Wait:     true,
	}

	var args []string
	args = append(args, diskArgs...)
	args = append(args, cloudInitArgs...)
	args = append(args, bridgeArgs...)

	if err := StartVM(config, args...); err != nil {
		fmt.Printf("Error launching VM: %v\n", err)
		os.Exit(1)
	}
}
