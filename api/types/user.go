package types_user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserStore interface {
	CreateUser(user RegisterUserPayload) (*User, error)
	GetUsers(query SearchUserQuery) ([]User, error)
	GetUserById(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsernameOrEmail(username string, email string) (*User, error)
	DeleteUserById(id string) error
	DeleteUserByUsername(username string) error
}

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginUserPayload struct {
	UsernameOrEmail string `json:"usernameOrEmail" validate:"required"`
	Password        string `json:"password"        validate:"required,min=6,max=130"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"  validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Username  string `json:"username"  validate:"required"`
	Password  string `json:"password"  validate:"required,min=6,max=130"`
}

type SearchUserQuery struct {
	Username string `json:"username"`
}

type UserJWTClaims struct {
	UserId    string `json:"user_id"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (c *UserJWTClaims) PopulateFromToken(claims jwt.MapClaims) error {
	c.UserId = claims["userId"].(string)
	c.ExpiresAt = int64(claims["expiresAt"].(float64))
	return nil
}
