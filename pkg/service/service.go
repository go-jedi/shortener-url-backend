package service

import (
	"github.com/rob-bender/shortener-url-backend/pkg/repository"
)

type TodoUrl interface {
	AddUrl(url string) (string, int, error)
}

type Service struct {
	TodoUrl
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUrl: NewUrlService(r.TodoUrl),
	}
}
