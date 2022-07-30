package config

import (
	"encoding/json"
	"os"
	"time"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type LFU struct {
	Size int           `json:"size"`
	TTL  time.Duration `json:"ttl_minute"`
}
type Redis struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Address string `json:"address"`
	LFU     LFU    `json:"lfu"`
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

func Load() (*Config, error) {
	b, err := os.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	config.Redis.LFU.TTL *= time.Minute
	return &config, nil
}
