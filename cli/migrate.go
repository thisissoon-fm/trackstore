package cli

import (
	"trackstore/db"
	"trackstore/log"

	"github.com/spf13/cobra"
)

var migrateSteps int

// Allows you to migrate up or down n number of steps
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if migrateSteps == 0 {
			cmd.Help()
			return
		}
		if err := db.Migrate(migrateSteps); err != nil {
			log.WithError(err).Debug("migrations not applied")
			return
		}
		log.Info("Applied migrations")
	},
}

// Runs up migrations
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run upgrade db migrations",
	Run: func(*cobra.Command, []string) {
		migrator, err := db.Migrator()
		if err != nil {
			log.WithError(err).Error("error applying migrations")
			return
		}
		defer migrator.Close()
		if err := migrator.Up(); err != nil {
			log.WithError(err).Debug("migrations not applied")
			return
		}
		log.Info("Applied migrations")
	},
}

// Runs down migrations
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Downgrade DB migrations by 1 migration step",
	Run: func(*cobra.Command, []string) {
		migrator, err := db.Migrator()
		if err != nil {
			log.WithError(err).Error("error applying migrations")
			return
		}
		defer migrator.Close()
		if err := migrator.Down(); err != nil {
			log.WithError(err).Debug("migrations not applied")
			return
		}
		log.Info("Applied migrations")
	},
}

func init() {
	migrateCmd.Flags().IntVar(&migrateSteps, "steps", 0, "Migration steps (+/-)")
	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd)
}
