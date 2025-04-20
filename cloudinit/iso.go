package cloudinit

import (
	"fmt"
	"os/exec"
)

func CreateCloudInitISO(outputFile string) error {
	cmd := exec.Command(
		"genisoimage",
		"-o", outputFile,
		"-volid", "cidata",
		"-joliet", "-rock",
		"ci-config/user-data",
		"ci-config/meta-data",
		"ci-config/network-config",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating ISO: %v\nOutput: %s", err, string(output))
	}

	return nil
}
