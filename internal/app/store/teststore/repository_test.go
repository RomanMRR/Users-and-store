package teststore_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.GetRepository().Create(model.TestUser(t)))
	assert.NotNil(t, u)
}

func TestRepository_FindBySurname(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	_, err := s.GetRepository().Find(u1.Surname, model.UserTable)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.GetRepository().Create(u1)
	u2, err := s.GetRepository().Find(u1.Surname, model.UserTable)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
