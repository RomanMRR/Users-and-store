package model

import (
	"testing"
	"time"
)

// Create and return Test User ...
func TestUser(t *testing.T) *User {
	var datetime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &User{
		Name:              "Roman",
		Surname:           "Romanov",
		Patronymic:        "Romanovich",
		Age:               17,
		Registration_date: datetime,
	}
}
