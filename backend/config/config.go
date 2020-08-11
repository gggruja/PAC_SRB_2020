package config

import (
	"encoding/json"
	"github.com/spf13/viper"
)

type Config struct {
	BindAddress    string
	DbDriver       string
	DbHost         string
	DbPort         string
	DbName         string
	DbUser         string
	DbPassword     string
}

// Default config for running the service locally
var Defaults = map[string]string{
	"BIND_ADDRESS":    ":9090",
	"DB_DRIVER":       "mysql",
	"DB_NAME":         "backend",
	"DB_USER":         "backend",
	"DB_PASSWORD":     "afPoLSv5s8Carbbw",
	"DB_HOST":         "backend-mariadb",
	"DB_PORT":         "3306",
}

func LoadConfig() (*Config, error) {
	configReader := viper.New()

	configReader.SetDefault("BIND_ADDRESS", Defaults["BIND_ADDRESS"])
	configReader.SetDefault("DB_DRIVER", Defaults["DB_DRIVER"])
	configReader.SetDefault("DB_NAME", Defaults["DB_NAME"])
	configReader.SetDefault("DB_USER", Defaults["DB_USER"])
	configReader.SetDefault("DB_PASSWORD", Defaults["DB_PASSWORD"])
	configReader.SetDefault("DB_HOST", Defaults["DB_HOST"])
	configReader.SetDefault("DB_PORT", Defaults["DB_PORT"])

	configReader.AutomaticEnv()

	configReader.SetConfigType("env")
	configReader.SetConfigName(".env")
	configReader.AddConfigPath(".")

	if err := configReader.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore, as it is optional
		} else {
			return nil, err
		}
	}

	config := Config{}

	config.BindAddress = configReader.GetString("BIND_ADDRESS")
	config.DbDriver = configReader.GetString("DB_DRIVER")
	config.DbHost = configReader.GetString("DB_HOST")
	config.DbPort = configReader.GetString("DB_PORT")
	config.DbName = configReader.GetString("DB_NAME")
	config.DbUser = configReader.GetString("DB_USER")
	config.DbPassword = configReader.GetString("DB_PASSWORD")

	return &config, nil
}

func (c *Config) String() string {
	s, _ := json.MarshalIndent(c, "", "\t")
	return string(s)
}
