package teststore

import (
	"errors"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
)

// Repository for testing
type Repository struct {
	store *Store
	users map[int]*model.User
}

// Create test user
func (r *Repository) Create(data store.Data) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if u, ok := data.(*model.User); ok {
		u.ID = len(r.users)
		r.users[u.ID] = u
		return nil
	}
	return errors.New("There is no suitable type")
}

// Find the required entry
func (r *Repository) Find(whatLookingFor string, tableName string) ([]store.Data, error) {
	var result []store.Data
	for _, u := range r.users {
		if u.Surname == whatLookingFor {
			result = append(result, u)
		}
	}
	if len(result) != 0 {
		return result, nil
	}

	return nil, store.ErrRecordNotFound

}

func (r *Repository) Delete(id int, tableName string) error {
	//TODO
	return nil
}

func (r *Repository) Update(data store.Data) error {
	//TODO
	return nil
}
