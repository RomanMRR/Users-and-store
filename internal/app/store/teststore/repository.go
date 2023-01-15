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
	if _, ok := r.users[id]; !ok {
		return errors.New("No rows in result set")
	} else {
		delete(r.users, id)

		return nil
	}

}

func (r *Repository) Update(data store.Data) error {
	if u, ok := data.(*model.UpdateUserInput); ok {
		if _, ok := r.users[*u.ID]; !ok {
			return errors.New("No rows in result set")
		}

		if u.Name != nil {
			r.users[*u.ID].Name = *u.Name
		}

		if u.Age != nil {
			r.users[*u.ID].Age = *u.Age
		}

		if u.Patronymic != nil {
			r.users[*u.ID].Patronymic = *u.Patronymic
		}

		if u.Registration_date != nil {
			r.users[*u.ID].Registration_date = *u.Registration_date
		}

		return nil
	}

	return errors.New("There is no suitable type")
}
