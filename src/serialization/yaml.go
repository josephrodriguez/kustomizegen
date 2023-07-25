package serialization

import (
	"os"

	"gopkg.in/yaml.v2"
)

func MarshalToYAML(data interface{}, filename string) error {
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
