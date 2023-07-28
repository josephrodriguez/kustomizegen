package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephrodriguez/kustomizegen/configuration"
	"github.com/josephrodriguez/kustomizegen/filesystem"
	"github.com/josephrodriguez/kustomizegen/serialization"
	template "github.com/josephrodriguez/kustomizegen/templates"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
)

func GenerateOverlaysCommand(cmd *cobra.Command, args []string) {
	root, _ := cmd.Flags().GetString("root")

	absoluteRootPath, err := filepath.Abs(root)
	if err != nil {
		log.Fatal("Error getting absolute path: %w", err)
	}

	config, err := configuration.ReadConfigurationFile(absoluteRootPath)
	if err != nil {
		log.Fatal("Error reading configuration: %w", err)
	}

	for _, overlay := range config.Overlays {
		overlayPath, err := getOverlayPath(absoluteRootPath, &overlay)
		if err != nil {
			log.Fatal("Error:", err)
		}

		resourcesPath, err := filepath.Rel(overlayPath, absoluteRootPath)
		if err != nil {
			fmt.Println("Error calculating relative path:", err)
			return
		}

		ctx := &types.KustomizegenContext{
			Namespace: overlay.Name,
		}

		namePrefix := interpolate(ctx, overlay.NamePrefix.Value, config.NamePrefix.Value)
		nameSuffix := interpolate(ctx, overlay.NameSuffix.Value, config.NameSuffix.Value)

		kustomization := types.PrototypeKustomization()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//Resources
		kustomization.Resources = append(kustomization.Resources, resourcesPath)

		//Namespace Transformer
		nsTransformer := types.NewNamespaceTransformer(overlay.Name, false)

		serializedTransformer, err := serialization.MarshalToYAML(nsTransformer)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		kustomization.Transformers = append(kustomization.Transformers, serializedTransformer)
		kustomization.NamePrefix = namePrefix
		kustomization.NameSuffix = nameSuffix

		kustomizationFileConfiguration := types.KustomizationFileConfiguration{
			NamePrefix: config.NamePrefix.Rules,
			NameSuffix: config.NameSuffix.Rules,
		}

		configurationPath := filepath.Join(overlayPath, "kustomizeconfig", "config.yaml")
		newKustomization, err := writeKustomizationConfigurationFile(kustomization, kustomizationFileConfiguration, configurationPath)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		//Output YAML file path with the Kustomization
		outputPath := filepath.Join(overlayPath, "kustomization.yaml")

		// Marshal the Kustomization instance to YAML and save it to a file
		if err := serialization.MarshalToYAMLFile(newKustomization, outputPath); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Created Kustomization overlay: %s\n", overlay.Name)
	}
}

func getConfig(configFilePath string) (*types.KustomizegenConfiguration, error) {
	// Check if the file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("File does not exist: %s", configFilePath)
	}

	// Read configuration from the file
	config, err := configuration.ReadConfigurationFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading configuration: %w", err)
	}

	return config, nil
}

func getOverlayPath(root string, overlay *types.KustomizegenOverlay) (string, error) {
	parent := filepath.Dir(root)
	overlayPath := filepath.Join(parent, "overlays", overlay.Name)

	// Create the overlay directory if it doesn't exist
	if err := filesystem.CreateFolder(overlayPath); err != nil {
		return "", fmt.Errorf("Error creating directory: %w", err)
	}

	return overlayPath, nil
}

func interpolate(ctx interface{}, templates ...string) string {
	// Access and use the list of parameters of type string here
	for _, tmpl := range templates {
		result, err := template.InterpolateTemplate(tmpl, ctx)
		if err != nil {
			fmt.Println("Error:", err)
			return ""
		}

		if result != "" {
			return result
		}
	}

	return ""
}

func writeKustomizationConfigurationFile(kustomization *types.Kustomization, kustomizationConfigFile types.KustomizationFileConfiguration, outputPath string) (types.Kustomization, error) {
	hasNoPrefixConfiguration := len(kustomizationConfigFile.NamePrefix) == 0 && len(kustomizationConfigFile.NameSuffix) == 0

	if hasNoPrefixConfiguration {
		return *kustomization, nil
	}

	if err := serialization.MarshalToYAMLFile(kustomizationConfigFile, outputPath); err != nil {
		return types.Kustomization{}, err
	}

	// Create a new instance of KustomizationFileConfiguration with updated configurations
	newKustomization := types.Kustomization{
		APIVersion:     kustomization.APIVersion,
		Kind:           kustomization.Kind,
		NamePrefix:     kustomization.NamePrefix,
		NameSuffix:     kustomization.NameSuffix,
		Resources:      kustomization.Resources,
		Configurations: append(kustomization.Configurations, outputPath),
	}

	return newKustomization, nil
}
