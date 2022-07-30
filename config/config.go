package config

import (
	"encoding/json"
	"os"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type Redis struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Address string `json:"address"`
}

type Cassandra struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	KeySpace string `json:"key_space"`
}

type Options struct {
	Schema string `json:"schema"`
	Prefix string `json:"Prefix"`
}

type Config struct {
	Server    Server    `json:"server"`
	Redis     Redis     `json:"redis"`
	Cassandra Cassandra `json:"cassandra"`
	Options   Options   `json:"options"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
