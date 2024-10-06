package store

import "github.com/Splucheviy/gopherSchoolLesson/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

