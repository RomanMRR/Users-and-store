package model

import (
	"testing"
	"time"
)

// Test Event ...
func TestEvent(t *testing.T) *User {
	var datetime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &User{
		Name:              "Roman",
		Surname:           "Romanov",
		Patronymic:        "Romanovich",
		Age:               17,
		Registration_date: datetime,
	}
}
