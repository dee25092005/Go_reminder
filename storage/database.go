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
		CREATE TABLE IF NOT EXISTS profiles(
			user_id INT PRIMARY KEY REFERENCES users(id),
			biography TEXT NOT NULL,
			github_username VARCHAR(100) NOT NULL
		)
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

func (ss *SQLStore) SaveProfile(userID int, bio string, github string) error {
	query := `
		INSERT INTO profiles (user_id, biography, github_username) VALUES ($1, $2, $3);
	`
	_, err := ss.db.Exec(query, userID, bio, github)
	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
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

func (ss *SQLStore) GetFullProfile(id int) (UserProfile, bool) {
	query := `
		SELECT u.id, u.name, p.biography,p.github_username 
		FROM users u
		INNER JOIN profiles p ON u.id = p.user_id
		WHERE u.id = $1;
		`

	var dto UserProfile
	err := ss.db.QueryRow(query, id).Scan(&dto.ID, &dto.Name, &dto.Biography, &dto.Github)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserProfile{}, false
		}
		fmt.Printf("failed to get user: %v\n", err)
		return UserProfile{}, false
	}

	return dto, true
}
