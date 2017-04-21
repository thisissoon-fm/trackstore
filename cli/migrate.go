package cli

import (
	"database/sql"

	"trackstore/config"
	"trackstore/log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database"
	"github.com/mattes/migrate/database/postgres"
	"github.com/spf13/cobra"
	"upper.io/db.v2/postgresql"

	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func dbDriver() (database.Driver, error) {
	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: config.DBName(),
		Host:     config.DBHost(),
		User:     config.DBUser(),
		Password: config.DBPass(),
	})
	if err != nil {
		return nil, err
	}
	db := sess.Driver().(*sql.DB)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func migrator() (*migrate.Migrate, error) {
	driver, err := dbDriver()
	if err != nil {
		return nil, err
	}
	return migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run database upgrade migrations",
	Run: func(*cobra.Command, []string) {
		migrator, err := migrator()
		if err != nil {
			log.WithError(err).Error("error loading migrator")
			return
		}
		defer migrator.Close()
		if err := migrator.Up(); err != nil {
			log.WithError(err).Error("errror running migrations")
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Run database downgrade migrations",
	Run: func(*cobra.Command, []string) {
		migrator, err := migrator()
		if err != nil {
			log.WithError(err).Error("error loading migrator")
			return
		}
		defer migrator.Close()
		if err := migrator.Down(); err != nil {
			log.WithError(err).Error("errror running migrations")
		}
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd)
}
