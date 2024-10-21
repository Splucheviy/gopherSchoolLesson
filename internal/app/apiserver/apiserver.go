package apiserver

import (
	"database/sql"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)
	return srv.router.Start(config.ServerAddr)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err

	}

	return db, nil
}
