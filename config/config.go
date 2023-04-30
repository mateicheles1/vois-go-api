package config

type Config struct {
	Server ConfigServer
}

type ConfigServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
