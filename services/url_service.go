package services

import (
	"github.com/Jaraxzus/url-shortener/db"
	"github.com/Jaraxzus/url-shortener/utils"
)

type URLService interface {
	ShortenURL(url string) (string, error)
	GetURL(code string) (string, error)
}

type URLServiceImpl struct {
	repo db.MongoDBRepository
}

func NewURLService(repo db.MongoDBRepository) URLService {
	return &URLServiceImpl{
		repo: repo,
	}
}

func (s *URLServiceImpl) ShortenURL(url string) (string, error) {
	startCode := utils.GenerateCode(url)
	return s.TryAddURL(startCode, url)
}

func (s *URLServiceImpl) TryAddURL(code string, url string) (string, error) {
	existingURL, err := s.GetURL(code)
	if err != nil {
		return "", err
	}

	if existingURL == "" {
		if err := s.repo.Set(code, url); err != nil {
			return "", err
		}
		return code, nil
	}

	if existingURL != url {
		newCode := utils.GenerateCode(url + code)
		return s.TryAddURL(newCode, url)
	}

	return code, nil
}

func (s *URLServiceImpl) GetURL(code string) (string, error) {
	url, err := s.repo.Get(code)
	if err != nil {
		return "", err
	}
	return url, nil
}
