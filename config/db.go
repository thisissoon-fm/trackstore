package config

import "github.com/spf13/viper"

const (
	db_host           = "db.host"
	db_name           = "db.name"
	db_user           = "db.user"
	db_pass           = "db.pass"
	db_migration_path = "db.migration_path"
)

func init() {
	viper.SetDefault(db_host, "localhost:5432")
	viper.SetDefault(db_name, "trackstore")
	viper.SetDefault(db_user, "postgres")
	viper.SetDefault(db_pass, "postgres")
	viper.SetDefault(db_migration_path, "migrations")
}

func DBHost() string {
	viper.BindEnv(db_host)
	return viper.GetString(db_host)
}

func DBName() string {
	viper.BindEnv(db_name)
	return viper.GetString(db_name)
}

func DBUser() string {
	viper.BindEnv(db_user)
	return viper.GetString(db_user)
}

func DBPass() string {
	viper.BindEnv(db_pass)
	return viper.GetString(db_pass)
}

func DBMigrationPath() string {
	viper.BindEnv(db_migration_path)
	return viper.GetString(db_migration_path)
}
