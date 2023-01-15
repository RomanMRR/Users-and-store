package sqlstore

import (
	"database/sql"
	"http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

// Store definition
type Store struct {
	db         *sql.DB
	repository *Repository
}

// New store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetRepository ...
func (s *Store) GetRepository() store.Repository {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &Repository{
		store: s,
	}

	return s.repository
}
