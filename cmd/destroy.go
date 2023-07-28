package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DestroyCommand(cmd *cobra.Command, args []string) {
	rootPath, _ := cmd.Flags().GetString("root")

	// Your implementation for the generate-build-command
	fmt.Println("Destroy executed!")
	fmt.Println("Root Path:", rootPath)
}
