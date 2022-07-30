package main

import (
	"fmt"
	"log"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/ctirouzh/tiny-url/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config file", err)
	}

	cassandra := storage.GetCassandraInstance(cfg.Cassandra)
	defer cassandra.Session.Close()
	if err != nil {
		fmt.Println("[main]<--", err)
	}
	fmt.Println("is cassandra session closed?", cassandra.Session.Closed())
	cache := storage.GetRedisCache(cfg.Redis)
	fmt.Println(cache)
}
