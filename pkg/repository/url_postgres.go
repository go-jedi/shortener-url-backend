package repository

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UrlPostgres struct {
	db *sqlx.DB
}

func NewUrlPostgres(db *sqlx.DB) *UrlPostgres {
	return &UrlPostgres{
		db: db,
	}
}

func (r *UrlPostgres) AddUrl(url string) (string, int, error) {
	var uid string
	err := r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	err = r.db.QueryRow("SELECT url_add($1, $2)", url, uid).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции url_add из базы данных, %s", err)
	}
	return uid, http.StatusOK, nil
}
