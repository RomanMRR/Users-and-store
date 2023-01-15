package sqlstore_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(model.UserTable)

	s := sqlstore.New(db)
	e := model.TestUser(t)
	assert.NoError(t, s.GetRepository().Create(model.TestUser(t)))
	assert.NotNil(t, e)
}

func TestRepository_FindBySurname(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(model.UserTable)

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.GetRepository().Create(u1)
	u2, err := s.GetRepository().Find(u1.Surname, model.UserTable)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(model.UserTable)

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	err := s.GetRepository().Delete(u1.ID, model.UserTable)
	assert.Error(t, err)
	s.GetRepository().Create(u1)
	err = s.GetRepository().Delete(u1.ID, model.UserTable)
	assert.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(model.UserTable)

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	newAge := int16(20)
	userUpdate := model.UpdateUserInput{
		ID:  &u1.ID,
		Age: &newAge,
	}
	err := s.GetRepository().Update(&userUpdate)
	assert.Error(t, err)
	s.GetRepository().Create(u1)
	err = s.GetRepository().Update(&userUpdate)
	assert.NoError(t, err)
}
