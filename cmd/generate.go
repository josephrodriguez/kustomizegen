package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateBuildCommand(cmd *cobra.Command, args []string) {
	rootPath, _ := cmd.Flags().GetString("root")
	outputPath, _ := cmd.Flags().GetString("output")

	// Your implementation for the generate-build-command
	fmt.Println("Generate-build-command executed!")
	fmt.Println("Root Path:", rootPath)
	fmt.Println("Output Path:", outputPath)
}
