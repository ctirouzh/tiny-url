package main

import (
	"fmt"
	"log"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/ctirouzh/tiny-url/storage"
)

func main() {

	cfg, err := config.Load("./config/config.json")
	if err != nil {
		log.Fatal("failed to load config file", err)
	}

	cassandra := storage.GetCassandraInstance(cfg.Cassandra.Host, cfg.Cassandra.KeySpace)
	defer cassandra.Session.Close()
	if err != nil {
		fmt.Println("[main]<--", err)
	}

	fmt.Println("is cassandra session closed?", cassandra.Session.Closed())
	cache := storage.GetRedisCache(cfg.Redis.Address)
	fmt.Println(cache)

	for i := 1; i < 5; i++ {
		storage.GetRedisCache(cfg.Redis.Address)
	}
}
