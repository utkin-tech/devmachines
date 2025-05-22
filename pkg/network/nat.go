package network

import (
	"errors"
	"fmt"
	"strings"

	"github.com/docker/go-connections/nat"
)

type Proto string

const (
	ProtoTcp = "tcp"
	ProtoUdp = "udp"
)

type Hostfwd struct {
	Proto     Proto
	Hostip    string
	Hostport  string
	Guestip   string
	Guestport string
}

func (h *Hostfwd) ToOption() string {
	return fmt.Sprintf(
		"hostfwd=%s:%s:%s-%s:%s",
		h.Proto,
		h.Hostip,
		h.Hostport,
		h.Guestip,
		h.Guestport,
	)
}

func SetupNAT(ports []string) ([]string, error) {
	options := []string{"user", "id=net0"}

	for _, port := range ports {
		hostfwd, err := portToHostfwd(port)
		if err != nil {
			return nil, err
		}

		option := hostfwd.ToOption()
		options = append(options, option)
	}

	fwdArg := strings.Join(options, ",")

	return []string{
		"-netdev", fwdArg,
		"-device", "virtio-net-pci,netdev=net0",
	}, nil
}

func portToHostfwd(port string) (*Hostfwd, error) {
	portMappings, err := nat.ParsePortSpec(port)
	if err != nil {
		return nil, err
	}

	if len(portMappings) != 1 {
		return nil, errors.New("expect single port mapping")
	}
	portMapping := portMappings[0]

	return &Hostfwd{
		Proto:     Proto(portMapping.Port.Proto()),
		Hostip:    portMapping.Binding.HostIP,
		Hostport:  portMapping.Binding.HostPort,
		Guestip:   "",
		Guestport: portMapping.Port.Port(),
	}, nil
}
