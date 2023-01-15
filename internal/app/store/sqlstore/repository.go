package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"strings"
)

// Repository devinition
type Repository struct {
	store *Store
}

// Create new data
func (r *Repository) Create(data store.Data) error {

	if err := data.Validate(); err != nil {
		return err
	}

	u, ok := data.(*model.User)
	if ok {
		query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age) VALUES ($1, $2, $3, $4) RETURNING id", model.UserTable)
		return r.store.db.QueryRow(
			query,
			u.Name,
			u.Surname,
			u.Patronymic,
			u.Age,
		).Scan(&u.ID)
	} else if s, ok := data.(*model.Shop); ok {
		query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age) VALUES ($1, $2, $3, $4) RETURNING id", model.ShopTable)
		return r.store.db.QueryRow(
			query,
			s.Name,
			s.Address,
			s.Working,
			s.Owner,
		).Scan(&s.ID)
	}

	return errors.New("There is no suitable type")

}

// Update data
func (r *Repository) Update(data store.Data) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if u, ok := data.(*model.UpdateUserInput); ok {
		if u.Name != nil {
			setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
			args = append(args, u.Name)
			argId++
		}

		if u.Age != nil {
			setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
			args = append(args, u.Age)
			argId++
		}

		if u.Patronymic != nil {
			setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
			args = append(args, u.Patronymic)
			argId++
		}

		if u.Registration_date != nil {
			setValues = append(setValues, fmt.Sprintf("registration_date=$%d", argId))
			args = append(args, u.Registration_date)
			argId++
		}

		setQuery := strings.Join(setValues, ", ")
		args = append(args, u.ID)
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d RETURNING id", model.UserTable, setQuery, argId)
		fmt.Println(query)
		fmt.Println(args...)
		return r.store.db.QueryRow(
			query, args...,
		).Scan(&u.ID)
	} else if shop, ok := data.(*model.UpdateShopInput); ok {
		if shop.Name != nil {
			setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
			args = append(args, u.Name)
			argId++
		}

		if shop.Address != nil {
			setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
			args = append(args, shop.Address)
			argId++
		}

		if shop.Working != nil {
			setValues = append(setValues, fmt.Sprintf("working=$%d", argId))
			args = append(args, shop.Working)
			argId++
		}

		if shop.Owner != nil {
			setValues = append(setValues, fmt.Sprintf("owner=$%d", argId))
			args = append(args, shop.Owner)
			argId++
		}

		setQuery := strings.Join(setValues, ", ")
		args = append(args, shop.ID)
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d RETURNING id", model.ShopTable, setQuery, argId)
		fmt.Println(query)
		fmt.Println(args...)
		return r.store.db.QueryRow(
			query, args...,
		).Scan(&shop.ID)
	}

	return errors.New("There is no suitable type")

}

// Delete records
func (r *Repository) Delete(id int, tableName string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", tableName)
	return r.store.db.QueryRow(
		query,
		id,
	).Scan(&id)
}

// Find data ...
func (r *Repository) Find(whatLookingFor string, tableName string) ([]store.Data, error) {

	if tableName == model.UserTable {
		u := &model.User{}
		query := fmt.Sprintf("SELECT * FROM %s WHERE surname = $1", tableName)
		rows, err := r.store.db.Query(
			query,
			whatLookingFor,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

		}

		var users []store.Data

		for rows.Next() {
			err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Age, &u.Registration_date)
			if err != nil {
				return nil, err
			}
			users = append(users, u)
		}

		return users, nil
	} else {
		s := &model.Shop{}
		query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableName)
		rows, err := r.store.db.Query(
			query,
			whatLookingFor,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

		}

		var shops []store.Data

		for rows.Next() {
			err := rows.Scan(&s.ID, &s.Name, &s.Address, &s.Working, &s.Owner)
			if err != nil {
				return nil, err
			}
			shops = append(shops, s)
		}

		return shops, nil
	}

}
