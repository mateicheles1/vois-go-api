package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s", c.Host, c.Port, c.User, c.Password, c.Database)
}

func LoadConfig(configFilePath string) (*Config, error) {

	file, err := os.Open(configFilePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
