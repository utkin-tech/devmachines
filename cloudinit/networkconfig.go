package cloudinit

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var DefaultEthernet = Ethernet{
	DHCP4: false,
	Nameservers: Nameserver{
		Addresses: []string{"8.8.8.8", "1.1.1.1"},
	},
}

type NetworkConfig struct {
	Version   int                 `yaml:"version"`
	Ethernets map[string]Ethernet `yaml:"ethernets"`
}

type Nameserver struct {
	Addresses []string `yaml:"addresses"`
}

type Ethernet struct {
	DHCP4       bool       `yaml:"dhcp4"`
	Addresses   []string   `yaml:"addresses"`
	Gateway4    string     `yaml:"gateway4"`
	Nameservers Nameserver `yaml:"nameservers"`
}

func NewNetworkConfig(ethernet *Ethernet) *NetworkConfig {
	return &NetworkConfig{
		Version: 2,
		Ethernets: map[string]Ethernet{
			"ens3": *ethernet,
		},
	}
}

func GenerateNetworkConfig(netplanConfig *NetworkConfig) (string, error) {
	yamlData, err := yaml.Marshal(&netplanConfig)
	if err != nil {
		return "", fmt.Errorf("error marshal YAML: %w", err)
	}

	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, "network-config")

	err = os.WriteFile(filePath, []byte(yamlData), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing file: %w", err)
	}

	return filePath, nil
}
