package main

import (
	"fmt"
	"log"
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

	_, err = disk.SetupDisk(user)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	_, err = cloudinit.SetupCloudInit(net, user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = network.SetupBridge(net)
	if err != nil {
		fmt.Printf("Error setting up network bridge: %v\n", err)
		return
	}

	// config := QEMUConfig{
	// 	MemoryMB:      2048,
	// 	CPUCores:      2,
	// 	DiskImagePath: "/blobs/disk.img",
	// 	SeedImagePath: "/blobs/seed.iso",
	// 	TapInterface:  "tap0",
	// 	Acceleration:  true,
	// 	Output:        os.Stdout,
	// 	Wait:          true,
	// }

	// if err := LaunchQEMUVM(config); err != nil {
	// 	fmt.Printf("Error launching VM: %v\n", err)
	// 	os.Exit(1)
	// }
}
