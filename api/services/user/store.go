package user

import (
	"database/sql"
	"fmt"

	types_user "megavault/api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types_user.RegisterUserPayload) (*types_user.User, error) {
	_, err := s.db.Exec(
		"INSERT INTO users (firstname, lastname, username, email, password) VALUES (?, ?, ?, ?, ?)",
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	u, err := s.GetUserByUsername(user.Username)
	if err != nil || u == nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) GetUsers(query types_user.SearchUserQuery) ([]types_user.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username LIKE '%?%'", query.Username)
	if err != nil {
		return nil, err
	}

	users := []types_user.User{}

	for rows.Next() {
		user, err := scanRow(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}

func (s *Store) GetUserById(id string) (*types_user.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(types_user.User)

	for rows.Next() {
		user, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (s *Store) GetUserByUsername(username string) (*types_user.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	user := new(types_user.User)

	for rows.Next() {
		user, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types_user.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types_user.User)

	for rows.Next() {
		user, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (s *Store) GetUserByUsernameOrEmail(username string, email string) (*types_user.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ? OR email = ?", username, email)
	if err != nil {
		return nil, err
	}

	user := new(types_user.User)

	for rows.Next() {
		user, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (s *Store) DeleteUserById(id string) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUserByUsername(username string) error {
	_, err := s.db.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		return err
	}

	return nil
}

func scanRow(rows *sql.Rows) (*types_user.User, error) {
	user := new(types_user.User)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
