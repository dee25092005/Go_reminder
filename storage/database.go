package storage

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type SQLStore struct {
	db *sql.DB
}

// postgresql://postgres:123@localhost:5432/test

func NewSQLStore(connStr string) (*SQLStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	store := &SQLStore{
		db: db,
	}

	err = store.createSchema()
	if err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	return store, nil
}

func (ss *SQLStore) createSchema() error {
	query :=
		`
		CREATE TABLE IF NOT EXISTS users(
			id INT PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
	`

	_, err := ss.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}

func (ss *SQLStore) Save(id int, name string) error {
	if name == "" {
		return ErrInvalidName
	}
	query := `
		INSERT INTO users (id, name) VALUES ($1, $2);
		`

	_, err := ss.db.Exec(query, id, name)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (ss *SQLStore) Get(id int) (User, bool) {
	query := `
		SELECT id, name FROM users WHERE id = $1;
	`
	var u User
	err := ss.db.QueryRow(query, id).Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, false
		}
		fmt.Printf("failed to get user: %v\n", err)
		return User{}, false

	}
	return u, true

}
