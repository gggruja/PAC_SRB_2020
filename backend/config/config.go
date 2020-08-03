package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BindAddress string
	BindDatabase string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	return &Config{
		BindAddress: os.Getenv("BIND_ADDRESS"),
		BindDatabase: os.Getenv("BIND_DATABASE"),
	}, nil
}

func (c *Config) String() string {
	s, _ := json.MarshalIndent(c, "", "\t")
	return string(s)
}
