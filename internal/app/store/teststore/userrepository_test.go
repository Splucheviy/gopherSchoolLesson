package teststore_test

import (
	"testing"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/model"
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}