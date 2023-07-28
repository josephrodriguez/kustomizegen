package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/josephrodriguez/kustomizegen/configuration"
	"github.com/josephrodriguez/kustomizegen/filesystem"
	"github.com/spf13/cobra"
)

func DestroyCommand(cmd *cobra.Command, args []string) {
	rootPath, _ := cmd.Flags().GetString("root")

	absoluteRootPath, err := filepath.Abs(rootPath)
	if err != nil {
		log.Fatal("Error getting absolute path: %w", err)
	}

	config, err := configuration.ReadConfigurationFile(absoluteRootPath)
	if err != nil {
		log.Fatal("Error reading configuration: %w", err)
	}

	parentPath := filepath.Dir(absoluteRootPath)

	for _, overlay := range config.Overlays {
		overlayPath := filepath.Join(parentPath, "overlays", overlay.Name)

		err := filesystem.DeleteDir(overlayPath)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Destroyed overlay:", overlay.Name)
	}
}
