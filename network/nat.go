package network

func CreateNatParams() []string {
	return []string{"-netdev", "user,id=net0,hostfwd=tcp::2222-:22"}
}
