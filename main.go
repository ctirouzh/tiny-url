package main

import (
	"fmt"
	"log"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/ctirouzh/tiny-url/middleware"
	"github.com/ctirouzh/tiny-url/service"
	"github.com/ctirouzh/tiny-url/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middleware.ErrorsMiddleware(gin.ErrorTypeAny))

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config file", err)
	}

	session := storage.GetCassandraInstance(cfg.Cassandra).Session
	defer session.Close()
	cache := storage.GetRedisCache(cfg.Redis)
	fmt.Println(cache)
	jwtService := service.NewJwtService(cfg.JWT.TTL, cfg.JWT.Secret, cfg.JWT.Issuer)
	fmt.Println(*jwtService)
}
