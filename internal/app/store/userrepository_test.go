package store_test

import (
	"testing"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/model"
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s, teardown := store.TestStore(t, "postgresql://postgres:password@127.0.0.1:5432/restapi_test?sslmode=disable")
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, "postgresql://postgres:password@127.0.0.1:5432/restapi_test?sslmode=disable")
	defer teardown("users")

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
