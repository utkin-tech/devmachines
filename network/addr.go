package network

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

func GetAddressesByInterface(ifaceName string) ([]string, error) {
	var addresses []string

	cmd := exec.Command("ip", "-4", "-j", "address", "show", ifaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run ip command: %v", err)
	}

	var interfaces []NetworkInterface
	if err := json.Unmarshal(output, &interfaces); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	if len(interfaces) != 1 {
		return nil, fmt.Errorf("expected single interface")
	}
	iface := interfaces[0]

	for _, addr := range iface.AddrInfo {
		if addr.Family == "inet" && addr.Scope == "global" {
			address := fmt.Sprintf("%s/%d", addr.Local, addr.Prefixlen)
			addresses = append(addresses, address)
		}
	}

	return addresses, nil
}
