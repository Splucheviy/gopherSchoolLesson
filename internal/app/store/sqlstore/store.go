package sqlstore

import (
	"database/sql"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store"
	_ "github.com/lib/pq"
)

// Store...
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// New store...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User...
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
