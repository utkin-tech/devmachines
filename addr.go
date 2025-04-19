package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type IPAddrInfo struct {
	Family    string `json:"family"`
	Local     string `json:"local"`
	Prefixlen int    `json:"prefixlen"`
	Scope     string `json:"scope"`
}

type NetworkInterface struct {
	Ifindex  int          `json:"ifindex"`
	Ifname   string       `json:"ifname"`
	Flags    []string     `json:"flags"`
	AddrInfo []IPAddrInfo `json:"addr_info"`
}

func GetIPv4ByInterface(ifaceName string) (string, error) {
	cmd := exec.Command("ip", "-4", "-j", "address")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run ip command: %v", err)
	}

	var interfaces []NetworkInterface
	if err := json.Unmarshal(output, &interfaces); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, iface := range interfaces {
		if iface.Ifname == ifaceName {
			for _, addr := range iface.AddrInfo {
				if addr.Family == "inet" && addr.Scope == "global" {
					return fmt.Sprintf("%s/%d", addr.Local, addr.Prefixlen), nil
				}
			}
			return "", fmt.Errorf("no IPv4 address with scope global found for interface %s", ifaceName)
		}
	}

	return "", fmt.Errorf("interface %s not found", ifaceName)
}
