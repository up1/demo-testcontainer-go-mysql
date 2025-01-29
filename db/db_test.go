package db_test

import (
	"context"
	"database/sql"
	"demo/db"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestDB(t *testing.T) {
	// Setup test container once for all tests
	ctx := context.Background()
	mysqlContainer, err := mysql.Run(ctx,
		"mysql:8.0.36",
		mysql.WithDatabase("demo-data"),
		mysql.WithUsername("user01"),
		mysql.WithPassword("pass01"),
		mysql.WithScripts(filepath.Join("testdata", "schema.sql")),
	)
	if err != nil {
		assert.Fail(t, "failed to start container")
	}

	connStr, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		assert.Fail(t, "failed to get connection string")
	}

	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		assert.Fail(t, "failed to open db connection")
	}
	defer _db.Close()

	testDB := db.NewDB(_db)

	t.Run("Success get user by ID", func(t *testing.T) {
		user, err := testDB.FindById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "Somkiat", user.Name)
	})

	t.Run("Not found user by ID", func(t *testing.T) {
		user, err := testDB.FindById(999)
		assert.Nil(t, err)
		assert.Nil(t, user)
	})
}
