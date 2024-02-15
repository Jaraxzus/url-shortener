package services

import (
	"github.com/Jaraxzus/url-shortener/db"
)

type RedisService interface {
	SaveValue(key string, value string) error
	GetValue(key string) (string, error)
}

// RedisService - сервис для работы с Redis
type RedisServiceImpl struct {
	repo db.RedisRepository
}

// NewRedisService создает новый экземпляр RedisService
func NewRedisService(repo db.RedisRepository) RedisService {
	return &RedisServiceImpl{
		repo: repo,
	}
}

// SaveValue сохраняет значение в Redis
func (s *RedisServiceImpl) SaveValue(key string, value string) error {
	return s.repo.Set(key, value)
}

// GetValue извлекает значение из Redis
func (s *RedisServiceImpl) GetValue(key string) (string, error) {
	return s.repo.Get(key)
}

// Close закрывает подключение к Redis
func (s *RedisServiceImpl) Close() {
	s.repo.Close()
}
