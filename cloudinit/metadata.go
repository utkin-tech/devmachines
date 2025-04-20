package cloudinit

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MetaData struct {
	InstanceID    string `yaml:"instance-id"`
	LocalHostname string `yaml:"local-hostname"`
}

func GenerateMetaData(instanceID, localHostname string) (string, error) {
	config := MetaData{
		InstanceID:    instanceID,
		LocalHostname: localHostname,
	}

	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		return "", fmt.Errorf("error marshal YAML: %w", err)
	}

	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, "meta-data")

	err = os.WriteFile(filePath, []byte(yamlData), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing file: %w", err)
	}

	return filePath, nil
}
