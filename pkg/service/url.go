package service

import (
	"github.com/rob-bender/shortener-url-backend/pkg/repository"
)

type UrlService struct {
	repo repository.TodoUrl
}

func NewUrlService(r repository.TodoUrl) *UrlService {
	return &UrlService{
		repo: r,
	}
}

func (s *UrlService) AddUrl(url string) (string, int, error) {
	return s.repo.AddUrl(url)
}

func (s *UrlService) GetUrl(uid string) (string, int, error) {
	return s.repo.GetUrl(uid)
}
