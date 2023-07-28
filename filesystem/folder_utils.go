package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFolder(path string, perm ...os.FileMode) error {
	// Set the default permission value if not provided
	defaultPerm := os.FileMode(0755)
	if len(perm) == 0 {
		perm = append(perm, defaultPerm)
	}

	err := os.MkdirAll(path, perm[0])
	if err != nil {
		return err
	}

	return nil
}

func ExistFolder(folderPath string) (bool, error) {
	// Check if the folder exists
	_, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("folder does not exist: %s", folderPath)
		}
		return false, err
	}
	return true, nil
}

// CreateTempDir creates a temporary directory and returns its path.
// The caller is responsible for cleaning up the temporary directory.
func CreateTempDir() (string, error) {
	tempDir, err := ioutil.TempDir("", "tempdir")
	if err != nil {
		return "", err
	}
	return tempDir, nil
}

// RemoveTempDir removes the temporary directory and its contents.
// If the directory does not exist, it returns nil.
func RemoveTempDir(tempDir string) error {
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		// Directory does not exist, return nil.
		return nil
	}

	return os.RemoveAll(tempDir)
}

// WriteFile writes data to a file with the given path and permissions.
// If the permissions argument is set to 0, it will use the default file permissions 0644.
func WriteFile(filePath string, data []byte, permissions os.FileMode) error {
	if permissions == 0 {
		permissions = 0644
	}

	err := ioutil.WriteFile(filePath, data, permissions)
	if err != nil {
		return err
	}
	return nil
}
