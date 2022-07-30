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
	Host    string        `json:"host"`
	Port    string        `json:"port"`
	Address string        `json:"address"`
	TTL     time.Duration `json:"ttl_hour"`
	LFU     LFU           `json:"lfu"`
}

type Cassandra struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	KeySpace string `json:"key_space"`
}

type JWT struct {
	TTL    time.Duration `json:"ttl_minute"`
	Secret string        `json:"secret"`
	Issuer string        `json:"issuer"`
}

type Options struct {
	Schema string `json:"schema"`
	Prefix string `json:"Prefix"`
}

type Config struct {
	Server    `json:"server"`
	Redis     `json:"redis"`
	Cassandra `json:"cassandra"`
	JWT       `json:"jwt"`
	Options   `json:"options"`
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
	config.JWT.TTL *= time.Minute
	config.Redis.TTL *= time.Hour
	config.Redis.LFU.TTL *= time.Minute
	return &config, nil
}
