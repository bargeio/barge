package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/bargeio/barge/pkg/config"
)

var bargeConfig *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "barge",
	Short: "Backup/restore for developers",
	Long: `Barge is a command-line tool to simplify typical backups and restore
operations during development.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		bargeConfig, err = config.GetConfig()
		if err != nil {
			// TODO give a better error message
			log.Fatal(err)
			return err
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.barge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
