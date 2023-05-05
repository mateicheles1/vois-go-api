package config

import (
	"encoding/json"
	"fmt"
	"gogin-api/logs"
	"os"
)

type Config struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"database"`
}
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s", c.Host, c.Port, c.User, c.Password, c.Database)
}

func NewConfig(configFilePath string) *Config {

	file, err := os.Open(configFilePath)

	if err != nil {
		logs.ErrorLogger.Error().Msgf("Error opening config file: %s", err)
		return nil
	}

	defer file.Close()

	var config Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		logs.ErrorLogger.Error().Msgf("Error unmarshaling json into config: %s", err)
		return nil
	}

	return &config
}
