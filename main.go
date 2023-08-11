package main

import (
	"fmt"

	"github.com/josephrodriguez/kustomizegen/cmd"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var rootCmd = &cobra.Command{Use: "kustomizegen"}

	rootCmd.AddCommand(
		generateOverlaysCmd(),
		generateBuildCmd(),
		destroyCmd(),
		buildCmd(),
		versionCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func generateOverlaysCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "generate-overlays",
		Short: "Create the Kustomization overlays",
		Run:   cmd.GenerateOverlaysCommand,
	}
	cmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	cmd.MarkFlagRequired("root")
	return cmd
}

func generateBuildCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "print-build-command",
		Short: "Generate the shell script with build commands for configured overlays",
		Run:   cmd.GenerateBuildCommand,
	}
	cmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	cmd.MarkFlagRequired("root")
	cmd.Flags().Bool("enable-helm", false, "Enable Helm")
	return cmd
}

func destroyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "destroy-overlays",
		Short: "Destroy the generated Kustomization overlays",
		Run:   cmd.DestroyOverlaysCommand,
	}
	cmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	cmd.MarkFlagRequired("root")
	return cmd
}

func buildCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "build",
		Short: "Build the project with specified options",
		Run:   cmd.BuildCommand,
	}
	cmd.Flags().StringP("root", "r", "", "Path to the Kustomization base folder")
	cmd.Flags().StringArrayP("overlays", "o", []string{"all"}, "Overlay's names")
	cmd.Flags().StringP("output-path", "p", "", "Output path for the build")
	cmd.Flags().StringP("output-template", "t", "", "Path template for the output file")
	cmd.Flags().BoolP("enable-helm", "H", false, "Enable Helm")
	cmd.Flags().StringP("load-restrictor", "L", "LoadRestrictionsRootOnly", "Load restrictor (LoadRestrictionsRootOnly or LoadRestrictionsNone)")
	cmd.Flags().StringArrayP("include", "i", []string{}, "Kubernetes resources to include in the build")
	cmd.Flags().StringArrayP("exclude", "e", []string{}, "Kubernetes resources to exclude from the build")
	return cmd
}

func versionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Display the compiled version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("KustomizeGen Version:", version)
		},
	}
	return cmd
}
