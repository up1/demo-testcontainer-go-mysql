package db

import "database/sql"

type User struct {
	ID   int
	Name string
}

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db: db}
}

func (d *DB) FindById(id int) (*User, error) {
	var user User
	err := d.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		// handle error
		recover()
		return nil, err
	}
	return &user, nil
}
