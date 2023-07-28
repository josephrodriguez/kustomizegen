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

func DeleteDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	return os.RemoveAll(path)
}
