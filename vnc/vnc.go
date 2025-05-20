package vnc

import "fmt"

const unixSocketPath = "/socks/vnc.sock"

func Setup(vnc string) []string {
	startServer()

	vncOption := fmt.Sprintf("unix:%s", unixSocketPath)

	args := []string{
		"-vga", "std",
		"-vnc", vncOption,
	}

	if len(vnc) == 0 {
		args = append(args, "-vnc", vnc)
	}

	return args
}
