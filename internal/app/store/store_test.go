package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://postgres:postgres@127.0.0.1:5432/restapi_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
