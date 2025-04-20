package main

import (
	"fmt"
	"log"
	"os"
)

const InterfaceName = "eth0"

func main() {
	ip, err := GetIPv4ByInterface(InterfaceName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("IPv4: %s\n", ip)

	gateway, err := GetDefaultGateway()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Default gateway: %s\n", gateway)

	err = CreateDiskImage(
		"/image/ubuntu.img",
		"/blobs/disk.img",
		"10G",
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// err := createISO("my_folder", "blobs/seed.iso")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ISO created successfully")

	// err := SetupNetworkBridge(
	// 	"br0",        // bridge name
	// 	"eth0",       // ethernet interface
	// 	"tap0",       // tap interface
	// 	"172.20.0.2/16", // IP to remove from eth0
	// )
	// if err != nil {
	// 	fmt.Printf("Error setting up network bridge: %v\n", err)
	// 	return
	// }

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
