package serial

func Setup() []string {
	return []string{
		"-chardev", "socket,id=serial0,path=/socks/serial.sock,server=on,wait=off",
		"-device", "isa-serial,chardev=serial0",
	}
}
