package cloudinit

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var DefaultUserData = UserData{
	Hostname:          "my-vm",
	ManageEtcHosts:    true,
	FQDN:              "my-vm",
	SSHAuthorizedKeys: []string{},
	Chpasswd: ChpasswdType{
		Expire: false,
	},
	Users: []string{
		"default",
	},
}

type ChpasswdType struct {
	Expire bool `yaml:"expire"`
}

type UserData struct {
	Hostname          string       `yaml:"hostname"`
	ManageEtcHosts    bool         `yaml:"manage_etc_hosts"`
	FQDN              string       `yaml:"fqdn"`
	User              string       `yaml:"user"`
	Password          string       `yaml:"password"`
	SSHAuthorizedKeys []string     `yaml:"ssh_authorized_keys"`
	Chpasswd          ChpasswdType `yaml:"chpasswd"`
	Users             []string     `yaml:"users"`
}

func GenerateUserData(userData *UserData) (string, error) {
	yamlData, err := yaml.Marshal(&userData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config: %v", err)
	}

	yamlContent := "#cloud-config\n" + string(yamlData)

	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, "user-data")

	err = os.WriteFile(filePath, []byte(yamlContent), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing file: %w", err)
	}

	return filePath, nil
}
