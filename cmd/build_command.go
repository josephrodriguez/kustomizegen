package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func BuildCommand(cmd *cobra.Command, args []string) {
	outputPath, _ := cmd.Flags().GetString("output-path")
	outputTemplate, _ := cmd.Flags().GetString("output-template")

	if outputPath != "" && outputTemplate != "" {
		panic("output template cannot be used when output path is specified")
	}

	fmt.Println("Build command executed successfully")
}
