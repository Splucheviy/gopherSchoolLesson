package teststore

import (
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/model"
	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
)

// Store...
type Store struct {
	UserRepository *UserRepository
}

// New store...
func New() *Store {
	return &Store{}
}

// User...
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.UserRepository
}
