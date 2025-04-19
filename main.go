package main

import (
	"fmt"
	"log"
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
}
