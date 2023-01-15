package teststore

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
)

// Store definition
type Store struct {
	repository *Repository
}

// New Store
func New() *Store {
	return &Store{}
}

// Get repostory
func (s *Store) GetRepository() store.Repository {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &Repository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.repository
}
