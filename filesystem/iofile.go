package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile writes data to a file with the given path and permissions.
// If the permissions argument is set to 0, it will use the default file permissions 0644.
func WriteFile(filePath string, data []byte, permissions ...os.FileMode) error {
	defaultPerm := os.FileMode(0644)
	if len(permissions) == 0 {
		permissions = append(permissions, defaultPerm)
	}

	// Get the directory path of the file
	dir := filepath.Dir(filePath)

	// Create the directories in the path if they don't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Write the file with the data and permissions
	err := ioutil.WriteFile(filePath, data, permissions[0])
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
