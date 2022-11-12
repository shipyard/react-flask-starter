package storage

import (
	"context"
	"database/sql"
	"net/http"
)

type search interface {
	Search(ctx context.Context, keys string) (string, error)
}

type Storage struct {
	search         search
	db             *sql.DB
	uploadendpoint string
	httpclient     *http.Client
}

func New(s search, db *sql.DB, ue string) *Storage {
	return &Storage{
		search:         s,
		db:             db,
		uploadendpoint: ue,
		httpclient:     &http.Client{},
	}
}
