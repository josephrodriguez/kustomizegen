package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"josephrodriguez.github.com/kustomizegen/src/configuration"
	"josephrodriguez.github.com/kustomizegen/src/filesystem"
	"josephrodriguez.github.com/kustomizegen/src/serialization"
	"josephrodriguez.github.com/kustomizegen/src/types"
)

func main() {
	configFilePath := flag.String("configuration", "config.yaml", "Path to the YAML configuration file")

	// Define a command-line flag for the root (folder path)
	var root string
	flag.StringVar(&root, "root", ".", "Root folder path")
	flag.Parse()

	// Validate if the root folder exists
	if err := filesystem.ExistFolder(root); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if the file exists
	if _, err := os.Stat(*configFilePath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist:", *configFilePath)
		return
	}

	// Read configuration from the file
	config, err := configuration.ReadConfigFromFile(*configFilePath)
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	// Get the absolute path from the relative path
	absolutePath, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	fmt.Println("Absolute Path:", absolutePath)

	parent := filepath.Dir(absolutePath)

	for _, namespace := range config.Namespaces {
		overlayPath := filepath.Join(parent, "overlays", namespace.Name)

		if err := filesystem.CreateFolder(overlayPath); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		// Calculate the relative path between targetDir and baseDir
		relativePath, err := filepath.Rel(overlayPath, absolutePath)
		if err != nil {
			fmt.Println("Error calculating relative path:", err)
			return
		}

		resources := []string{
			relativePath,
		}

		kustomization := types.NewKustomization(resources)
		kustomizationFile := filepath.Join(overlayPath, "kustomization.yaml")

		// Marshal the Kustomization instance to YAML and save it to a file
		if err := serialization.MarshalToYAML(kustomization, kustomizationFile); err != nil {
			fmt.Println("Error saving YAML:", err)
			return
		}

		fmt.Println("Kustomization saved to kustomization.yaml")
	}
}
