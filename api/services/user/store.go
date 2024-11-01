package user

import (
	"database/sql"
	"fmt"

	"megavault/api/types/user"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types_user.RegisterUserPayload) (*types_user.User, error) {
	rowId := ""
	err := s.db.QueryRow(
		"INSERT INTO users (firstname,lastname,username,email,password) VALUES ($1,$2,$3,$4,$5) RETURNING id;",
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&rowId)
	if err != nil {
		return nil, err
	}

	u, err := s.GetUserById(rowId)
	if err != nil || u == nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) GetUsers(query types_user.SearchUserQuery) ([]types_user.User, error) {
	rows, err := s.db.Query(
		"SELECT * FROM users WHERE username LIKE $1;",
		fmt.Sprintf("%%%s%%", query.Username),
	)
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
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1;", id)
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
	rows, err := s.db.Query("SELECT * FROM users WHERE username = $1;", username)
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
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1;", email)
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
	rows, err := s.db.Query(
		"SELECT * FROM users WHERE username = $1 OR email = $2;",
		username,
		email,
	)
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
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1;", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUserByUsername(username string) error {
	_, err := s.db.Exec("DELETE FROM users WHERE username = $1;", username)
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
