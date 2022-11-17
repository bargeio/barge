package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export [flags] [asset]",
	Short: "Export database or files to the specified destination",
	Long:  `Export the specified assets to the specified destination.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
		fmt.Printf("Arguments: %v\n", args)
		fmt.Printf("Configuration: %v\n", bargeConfig)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
