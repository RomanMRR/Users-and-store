package model_test

import (
	"http-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestEvent(t)
			},
			isValid: true,
		},
		{
			name: "Big age",
			u: func() *model.User {
				u := model.TestEvent(t)
				u.Age = 396
				return u
			},
			isValid: false,
		},
		{
			name: "empty name",
			u: func() *model.User {
				u := model.TestEvent(t)
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short name",
			u: func() *model.User {
				e := model.TestEvent(t)
				e.Name = "s"
				return e
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
