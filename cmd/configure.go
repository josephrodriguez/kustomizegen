package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephrodriguez/kustomizegen/configuration"
	"github.com/josephrodriguez/kustomizegen/filesystem"
	"github.com/josephrodriguez/kustomizegen/serialization"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
)

func Configure(cmd *cobra.Command, args []string) {
	rootPath, _ := cmd.Flags().GetString("root")
	configFilePath, _ := cmd.Flags().GetString("configuration")

	// Call the private getConfig function
	config, err := getConfig(configFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the absolute path from the relative path
	absolutePath, err := filepath.Abs(rootPath)
	if err != nil {
		log.Fatal("Error getting absolute path: %w", err)
	}

	for _, namespace := range config.Namespaces {
		overlayPath, err := getOverlayPath(absolutePath, &namespace)
		if err != nil {
			log.Fatal("Error:", err)
		}

		resourcesPath, err := filepath.Rel(overlayPath, absolutePath)
		if err != nil {
			fmt.Println("Error calculating relative path:", err)
			return
		}

		kustomization, err := generateKustomization(namespace.Name, resourcesPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//Output YAML file path with the Kustomization
		outputPath := filepath.Join(overlayPath, "kustomization.yaml")

		// Marshal the Kustomization instance to YAML and save it to a file
		if err := serialization.MarshalToYAMLFile(kustomization, outputPath); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Kustomization saved for namespace: %s\n", namespace.Name)
	}
}

func getConfig(configFilePath string) (*types.KustomizationConfig, error) {
	// Check if the file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("File does not exist: %s", configFilePath)
	}

	// Read configuration from the file
	config, err := configuration.ReadConfigFromFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading configuration: %w", err)
	}

	return config, nil
}

func generateKustomization(namespace string, resourcesPath string) (*types.Kustomization, error) {
	nsTransformer := types.NewNamespaceTransformer(namespace, false)

	serializedTransformer, err := serialization.MarshalToYAML(nsTransformer)
	if err != nil {
		return nil, err
	}

	kustomization := &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Resources:  []string{resourcesPath},
		Transformers: []string{
			serializedTransformer,
		},
	}

	return kustomization, nil
}

func getOverlayPath(root string, namespace *types.KustomizationNamespace) (string, error) {
	parent := filepath.Dir(root)
	overlayPath := filepath.Join(parent, "overlays", namespace.Name)

	// Create the overlay directory if it doesn't exist
	if err := filesystem.CreateFolder(overlayPath); err != nil {
		return "", fmt.Errorf("Error creating directory: %w", err)
	}

	return overlayPath, nil
}
