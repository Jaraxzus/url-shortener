package main

import (
	"log"
	"time"

	"github.com/Jaraxzus/url-shortener/config"
	"github.com/Jaraxzus/url-shortener/db"
	"github.com/Jaraxzus/url-shortener/handlers"
	"github.com/Jaraxzus/url-shortener/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	log.Println("MongoDB:", cfg.MongoDB.URI)
	repo, err := db.NewMongoDBRepository(cfg.MongoDB.URI, cfg.MongoDB.DBName, cfg.MongoDB.CollName)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	redisRepo, err := db.NewRedisRepository(cfg.Redis.URI, time.Duration(cfg.Redis.CACHE_TIMEOUT)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	redisService := services.NewRedisService(*redisRepo)
	urlService := services.NewURLService(*repo)
	route := gin.Default()
	// Добавляем сервис в контекст Gin
	route.Use(func(c *gin.Context) {
		c.Set("urlService", urlService)
		c.Set("redisService", redisService)
		c.Next()
	})
	route.GET("/a/", handlers.ShortenHandler)
	route.GET("/s/:code", handlers.RedirectHandler)
	route.Run(":8080")
}
