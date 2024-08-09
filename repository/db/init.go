package db

import (
	"database/sql"

	repo "github.com/meivaldi/billing-engine/repository"
)

type RepositoryDB struct {
	db *sql.DB
}

func New(db *sql.DB) (repo.Repository, error) {
	return &RepositoryDB{
		db: db,
	}, nil
}
