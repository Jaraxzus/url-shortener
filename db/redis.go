package db

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisRepository предоставляет методы для работы с Redis
type RedisRepository struct {
	client   *redis.Client
	duration time.Duration
}

// NewRedisRepository создает новый экземпляр RedisRepository
func NewRedisRepository(addr string, duration time.Duration) (*RedisRepository, error) {
	opt, err := redis.ParseURL(addr)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisRepository{
		client:   client,
		duration: duration,
	}, nil
}

// Set сохраняет значение в Redis с временем жизни
func (r *RedisRepository) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(context.Background(), key, data, r.duration).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	data, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	r.client.Expire(context.Background(), key, r.duration)

	log.Println(data)
	return strings.Trim(data, `"`), nil
}

// Close закрывает подключение к Redis
func (r *RedisRepository) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
