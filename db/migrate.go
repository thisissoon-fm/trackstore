package db

import (
	"database/sql"
	"fmt"
	"strings"

	"trackstore/config"
	"trackstore/log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"upper.io/db.v3/postgresql"

	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

type MigrateLogger struct{}

func (*MigrateLogger) Printf(format string, v ...interface{}) {
	format = strings.Replace(format, "\n", "", -1)
	log.Debug(format, v...)
}

func (*MigrateLogger) Verbose() bool {
	return false
}

// Returns a DB Migrator
func Migrator() (*migrate.Migrate, error) {
	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: config.DBName(),
		Host:     config.DBHost(),
		User:     config.DBUser(),
		Password: config.DBPass(),
	})
	if err != nil {
		return nil, err
	}
	sqlDB := sess.Driver().(*sql.DB)
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("file://%s", config.DBMigrationPath())
	migrator, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		return nil, err
	}
	migrator.Log = &MigrateLogger{}
	return migrator, nil
}

// Migrate the database up or down a number of steps positive
// numbers go up, negative numbers go down
// Takes path to migrations directory and the number of steps to apply
func Migrate(steps int) error {
	migrator, err := Migrator()
	if err != nil {
		return err
	}
	defer migrator.Close()
	migrator.Log = &MigrateLogger{}
	return migrator.Steps(steps)
}
