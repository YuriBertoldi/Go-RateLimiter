package main

import (
	"log"
	"net/http"

	"github.com/YuriBertoldi/Go-RateLimiter/internal/bd"
	"github.com/YuriBertoldi/Go-RateLimiter/internal/config"
	"github.com/YuriBertoldi/Go-RateLimiter/internal/limiter"
	"github.com/YuriBertoldi/Go-RateLimiter/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error reading configuration file: %v", err)
	}

	store, err := bd.NewRedisClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	if err != nil {
		log.Fatalf("error connecting to Redis: %v", err)
	}

	limiter := limiter.NewRateLimit(cfg, store)

	router := chi.NewRouter()

	router.Use(middleware.RateLimiterMiddleware(limiter))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	})

	http.ListenAndServe(":8080", router)
}
