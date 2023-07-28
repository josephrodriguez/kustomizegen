package main

import (
	"github.com/josephrodriguez/kustomizegen/cmd"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var rootCmd = &cobra.Command{Use: "kustomizegen"}

	var configureCmd = &cobra.Command{
		Use:   "generate-overlays",
		Short: "Create the Kustomization overlays",
		Run:   cmd.GenerateOverlaysCommand,
	}
	configureCmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	configureCmd.MarkFlagRequired("root")

	var generateBuildCmd = &cobra.Command{
		Use:   "generate-build-command",
		Short: "Generate the shell script with build commands for configured overlays",
		Run:   cmd.GenerateBuildCommand,
	}
	generateBuildCmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	generateBuildCmd.Flags().StringP("output", "o", "", "Path to the output shell file")
	generateBuildCmd.MarkFlagRequired("root")
	generateBuildCmd.MarkFlagRequired("output")

	var destroyCmd = &cobra.Command{
		Use:   "destroy-overlays",
		Short: "Destroy the generated Kustomization overlays",
		Run:   cmd.DestroyCommand,
	}
	destroyCmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	destroyCmd.MarkFlagRequired("root")

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the compiled version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("KustomizeGen Version:", version)
		},
	}

	rootCmd.AddCommand(configureCmd, generateBuildCmd, destroyCmd, versionCmd)
	rootCmd.Execute()
}
