package dbrepo

import (
	"database/sql"
	"github.com/ory/dockertest/v3"
	"os"
	"testing"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB

//var testRepo repository.DatabaseRepo

func TestMain(m *testing.M) {
	// run tests
	code := m.Run()

	os.Exit(code)
}
