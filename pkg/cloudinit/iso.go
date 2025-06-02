package cloudinit

import (
	"fmt"
	"os/exec"

	"github.com/utkin-tech/devmachines/pkg/utils"
)

const CloudInitIsoPath = "/disks/cloudinit.iso"

type Addr interface {
	CIDR() string
}

type Network[T Addr] interface {
	Addresses() []T
	Gateway() string
}

type User interface {
	User() string
	Password() string
	SSHKeys() []string
}

func SetupCloudInit[T Addr](network Network[T], user User) ([]string, error) {
	if err := CreateISO(CloudInitIsoPath, network, user); err != nil {
		return nil, err
	}

	return []string{
		"-drive", fmt.Sprintf("file=%s,format=raw,if=virtio", CloudInitIsoPath),
	}, nil
}

func CreateISO[T Addr](outputFile string, network Network[T], user User) error {
	instanceID, err := utils.RandomHex(20)
	if err != nil {
		return err
	}

	metaDataPath, err := GenerateMetaData(&MetaData{
		InstanceID:    instanceID,
		LocalHostname: "my-vm",
	})
	if err != nil {
		return err
	}

	userData := DefaultUserData
	userData.User = user.User()
	userData.Password = user.Password()
	userData.SSHAuthorizedKeys = user.SSHKeys()
	userDataPath, err := GenerateUserData(&userData)
	if err != nil {
		return err
	}

	ethernet := DefaultEthernet
	ethernet.Gateway4 = network.Gateway()

	var addresses []string
	for _, addr := range network.Addresses() {
		addresses = append(addresses, addr.CIDR())
	}
	ethernet.Addresses = addresses

	networkConfig := NewNetworkConfig(&ethernet)
	networkConfigPath, err := GenerateNetworkConfig(networkConfig)
	if err != nil {
		return err
	}

	args := []string{
		"-o", outputFile,
		"-volid", "cidata",
		"-joliet", "-rock",
		metaDataPath,
		userDataPath,
	}

	addNetwork := false
	if addNetwork {
		args = append(args, networkConfigPath)
	}

	cmd := exec.Command(
		"genisoimage",
		args...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating ISO: %v\nOutput: %s", err, string(output))
	}

	return nil
}
