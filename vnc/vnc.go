package vnc

func Setup(vnc string) []string {
	if len(vnc) == 0 {
		return nil
	}

	return []string{
		"-vga", "std",
		"-vnc", vnc,
	}
}
