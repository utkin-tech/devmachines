package serial

import "fmt"

const unixSocketPath = "/socks/serial.sock"

func Setup() []string {
	startServer()

	chardevOption := fmt.Sprintf("socket,id=serial0,path=%s,server=on,wait=off", unixSocketPath)

	return []string{
		"-chardev", chardevOption,
		"-device", "isa-serial,chardev=serial0",
	}
}
