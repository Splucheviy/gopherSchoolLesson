package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

// TestStore...
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Error connection to database: %v", err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
