package model_test

import (
	"testing"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/model"
	"github.com/stretchr/testify/assert"
)

// TestUser_Validate...
func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isvalid bool
	}{
		{
			name: "Valid user",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isvalid: true,
		},
		{
			name: "with encryption password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"
				
				return u
			},
			isvalid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isvalid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid email"

				return u
			},
			isvalid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isvalid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"

				return u
			},
			isvalid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isvalid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

// TestUser_BeforeCreate...
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
