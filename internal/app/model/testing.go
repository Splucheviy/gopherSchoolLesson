package model

import "testing"

// TestUser...
func TestUser(t *testing.T) *User {
	return &User{
		Email: "test@example.com",
		Password: "password",
	}
}
