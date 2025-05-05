package serial

func Setup() []string {
	return []string{
		"-serial", "unix:/socks/serial.sock,server,nowait",
	}
}
