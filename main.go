package main

import (
	"fmt"
	"log"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/ctirouzh/tiny-url/controller"
	"github.com/ctirouzh/tiny-url/middleware"
	"github.com/ctirouzh/tiny-url/repo"
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
	cacheRepo := repo.NewCacheRepository(cache, cfg.Redis)
	urlRepo := repo.NewURLRepository(session, cacheRepo)
	urlService := service.NewUrlService(urlRepo)
	urlController := controller.NewURLController(urlService)

	jwtService := service.NewJwtService(cfg.JWT.TTL, cfg.JWT.Secret, cfg.JWT.Issuer)
	fmt.Println("[main]--> jwtService:", *jwtService)

	r.GET("/:hash", urlController.RedirectURLByHash)

	r.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
