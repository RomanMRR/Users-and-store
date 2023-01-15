package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

const Layout = "2006-01-02"
const UserTable = "users"

type User struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Patronymic        string    `json:"patronymic"`
	Age               int16     `json:"age"`
	Registration_date time.Time `json:"Registration_date"`
}

type UpdateUserInput struct {
	ID                *int       `json:"id"`
	Name              *string    `json:"name"`
	Surname           *string    `json:"surname"`
	Patronymic        *string    `json:"patronymic"`
	Age               *int16     `json:"age"`
	Registration_date *time.Time `json:"Registration_date"`
}

// Validate model user ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&u.Age, validation.Required, validation.Min(4), validation.Max(200)),
		validation.Field(&u.Surname, validation.Required),
	)
}

// Validate incoming request with...
func (i *UpdateUserInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Patronymic == nil && i.Age == nil && i.Registration_date == nil {
		return errors.New("User update structure has no values")
	}

	return nil
}
