package network

import "fmt"

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

func SetupNAT(fwd Hostfwd) ([]string, error) {
	fwdArg := fmt.Sprintf(
		"user,id=net0,hostfwd=%s:%s:%s-%s:%s",
		fwd.Proto,
		fwd.Hostip,
		fwd.Hostport,
		fwd.Guestip,
		fwd.Guestport,
	)

	return []string{
		"-netdev", fwdArg,
		"-device", "virtio-net-pci,netdev=net0",
	}, nil
}
