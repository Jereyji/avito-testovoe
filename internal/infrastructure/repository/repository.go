package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Repository struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}