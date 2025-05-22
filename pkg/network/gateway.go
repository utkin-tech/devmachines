package network

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Route struct {
	Dst      string `json:"dst"`
	Gateway  string `json:"gateway"`
	Dev      string `json:"dev"`
	Protocol string `json:"protocol"`
}

func GetDefaultGateway() (string, error) {
	cmd := exec.Command("ip", "-j", "route", "show", "default")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run ip command: %v", err)
	}

	var routes []Route
	err = json.Unmarshal(output, &routes)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, route := range routes {
		if route.Dst == "default" && route.Gateway != "" {
			return route.Gateway, nil
		}
	}

	return "", fmt.Errorf("no default gateway found")
}
