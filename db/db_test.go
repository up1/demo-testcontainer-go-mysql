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

// Testing with testcontainers + mysql

func TestSuccessWithGetDataByID(t *testing.T) {
	// Start a new container of MySQL Database
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

	// Act
	_db, err := sql.Open("mysql", connStr)
	testDB := db.NewDB(_db)
	user, err := testDB.FindById(1)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "Somkiat", user.Name)
}
