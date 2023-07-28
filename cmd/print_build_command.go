package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/josephrodriguez/kustomizegen/configuration"
	tmpl "github.com/josephrodriguez/kustomizegen/templates"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
)

func GenerateBuildCommand(cmd *cobra.Command, args []string) {
	rootPath, _ := cmd.Flags().GetString("root")
	enableHelm, _ := cmd.Flags().GetBool("enable-helm")

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

		ctx := types.KustomizegenBuildCommandContext{
			Path:       overlayPath,
			EnableHelm: enableHelm,
		}

		template := `kustomize build {{if .EnableHelm}}--enable-helm{{end}} {{.Path}}`

		result, err := tmpl.InterpolateTemplate(template, ctx)
		if err != nil {
			log.Fatal("Error:", err)
		}

		fmt.Println(result)
	}
}
