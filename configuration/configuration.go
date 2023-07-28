package configuration

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/josephrodriguez/kustomizegen/types"
	"gopkg.in/yaml.v2"
)

func ReadConfigurationFile(root string) (*types.KustomizegenConfiguration, error) {
	// Call the private getConfig function
	configFilePath := filepath.Join(root, "kustomizegen.yaml")

	// Open the YAML file
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal("Error getting config path: %w", err)
	}

	// Initialize a new instance of the struct
	var config types.KustomizegenConfiguration

	// Unmarshal the YAML content into the struct
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
