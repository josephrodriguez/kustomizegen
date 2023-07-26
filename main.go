package main

import (
	"github.com/josephrodriguez/kustomizegen/cmd"
	"github.com/spf13/cobra"
)

var version string // version will be set during the build using ldflags

func main() {
	var rootCmd = &cobra.Command{Use: "kustomizegen"}

	var configureCmd = &cobra.Command{
		Use:   "configure",
		Short: "Create the kustomization files",
		Run:   cmd.Configure,
	}
	configureCmd.Flags().StringP("namespace", "n", "", "Namespace value")
	configureCmd.Flags().BoolP("unsetOnly", "u", false, "UnsetOnly value")

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the compiled version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("KustomizeGen Version", version)
		},
	}

	rootCmd.AddCommand(configureCmd, versionCmd)
	rootCmd.Execute()
}
