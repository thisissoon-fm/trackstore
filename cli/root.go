package cli

import (
	"trackstore/app"
	"trackstore/config"
	"trackstore/log"

	"github.com/spf13/cobra"
)

// Custom configuration file path
var configPath string

// Root CLI Command
var rootCmd = &cobra.Command{
	Use:   "trackstore",
	Short: "Provides a HTTP API to 3rd Party Track Data",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Bind log-level flag to config to allow log level to be set from
		// the CLI Flag
		config.LogLevelFlag(cmd.PersistentFlags().Lookup("log-level"))
		// Load custom configuration file if path given
		if configPath != "" {
			if err := config.Read(configPath); err != nil {
				log.WithError(err).Error("error reading configuration")
			}
		}
		// Run Application Pre Run
		app.PreRun()
	},
	Run: func(*cobra.Command, []string) {
		if err := app.Run(); err != nil {
			log.WithError(err).Error("application run error")
		}
	},
}

// Initialiser
func init() {
	// Add CLI Flags to Root Command
	rootCmd.PersistentFlags().StringVarP(
		&configPath,
		"config",
		"c",
		"",
		"Absolute path to configuration file")
	rootCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Log Level (debug,info,warn,error)")
	// Add Sub Commands
	rootCmd.AddCommand(versionCmd)
}

// Execute CLI
func Execute() error {
	return rootCmd.Execute()
}
