package config

import (
	"os"
	"strings"

	"trackstore/log"

	"github.com/spf13/viper"
)

// Configuration defaults
func init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/sfm/trackstore")
	viper.AddConfigPath("$HOME/.config/sfm/trackstore")
	viper.SetEnvPrefix("TRACKSTORE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Warn("error reading configuration, using default values")
	}
}

// Load custom configuration file
func Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return viper.ReadConfig(f)
}
