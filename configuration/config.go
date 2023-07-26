package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"josephrodriguez.github.com/kustomizegen/src/types"
)

func ReadConfigFromFile(filePath string) (*types.KustomizationConfig, error) {
	// Open the YAML file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Initialize a new instance of the struct
	var config types.KustomizationConfig

	// Unmarshal the YAML content into the struct
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
