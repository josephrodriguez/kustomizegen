package serialization

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func MarshalToYAMLFile(data interface{}, filename string) error {
	// Get the directory path of the file
	dir := filepath.Dir(filename)

	// Create the directory and all parent directories if they don't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Open the file for writing with create flag if it doesn't exist
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a YAML encoder to write data to the file
	encoder := yaml.NewEncoder(file)

	// Encode the data (write YAML data to the file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func MarshalToYAML(data interface{}) (string, error) {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}
