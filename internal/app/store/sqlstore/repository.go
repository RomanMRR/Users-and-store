package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"strings"
)

// Repository ...
type Repository struct {
	store *Store
}

// Create ...
func (r *Repository) Create(data store.Data) error {

	if err := data.Validate(); err != nil {
		return err
	}

	u, ok := data.(*model.User)
	if ok {
		return r.store.db.QueryRow(
			"INSERT INTO users (name, surname, patronymic, age) VALUES ($1, $2, $3, $4) RETURNING id",
			u.Name,
			u.Surname,
			u.Patronymic,
			u.Age,
		).Scan(&u.ID)
	} else if s, ok := data.(*model.Shop); ok {
		return r.store.db.QueryRow(
			"INSERT INTO shops (name, address, working, owner) VALUES ($1, $2, $3, $4) RETURNING id",
			s.Name,
			s.Address,
			s.Working,
			s.Owner,
		).Scan(&s.ID)
	}

	return errors.New("There is no suitable type")

}

// Update ...
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
			setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
			args = append(args, u.Patronymic)
			argId++
		}

		setQuery := strings.Join(setValues, ", ")
		args = append(args, u.ID)
		query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d RETURNING id", setQuery, argId)
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
		query := fmt.Sprintf("UPDATE shops SET %s WHERE id = $%d RETURNING id", setQuery, argId)
		fmt.Println(query)
		fmt.Println(args...)
		return r.store.db.QueryRow(
			query, args...,
		).Scan(&shop.ID)
	}

	return errors.New("")

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

	if tableName == "users" {
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

// // Find for month ...
// func (r *EventRepository) FindForMonth(month, year string, user_id int) ([]model.Event, error) {
// 	e := &model.Event{}
// 	rows, err := r.store.db.Query(
// 		"SELECT id, user_id, name, date FROM events WHERE extract(month from date) = $1 AND extract(year from date) = $2 and user_id = $3",
// 		month,
// 		year,
// 		user_id,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, store.ErrRecordNotFound
// 		}

// 	}

// 	var events []model.Event

// 	for rows.Next() {
// 		err := rows.Scan(&e.ID, &e.User_id, &e.Name, &e.Date)
// 		if err != nil {
// 			return nil, err
// 		}
// 		events = append(events, *e)
// 	}

// 	return events, nil
// }

// // Find for week ...
// func (r *EventRepository) FindForWeek(week, year string, user_id int) ([]model.Event, error) {
// 	e := &model.Event{}
// 	rows, err := r.store.db.Query(
// 		"SELECT id, user_id, name, date FROM events WHERE extract(week from date) = $1 AND extract(year from date) = $2 and user_id = $3",
// 		week,
// 		year,
// 		user_id,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, store.ErrRecordNotFound
// 		}

// 	}

// 	var events []model.Event

// 	for rows.Next() {
// 		err := rows.Scan(&e.ID, &e.User_id, &e.Name, &e.Date)
// 		if err != nil {
// 			return nil, err
// 		}
// 		events = append(events, *e)
// 	}

// 	return events, nil
// }
