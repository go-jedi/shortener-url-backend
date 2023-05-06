package repository

import (
	"github.com/jmoiron/sqlx"
)

type TodoUrl interface {
	AddUrl(url string) (string, int, error)
}

type Repository struct {
	TodoUrl
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUrl: NewUrlPostgres(db),
	}
}
