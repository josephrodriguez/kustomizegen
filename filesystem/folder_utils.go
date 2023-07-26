package filesystem

import (
	"fmt"
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

func ExistFolder(folderPath string) error {
	// Check if the folder exists
	_, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("folder does not exist: %s", folderPath)
		}
		return err
	}
	return nil
}
