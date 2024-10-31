package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TestUserJWTClaims struct {
	UserId    string `json:"user_id"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (c *TestUserJWTClaims) PopulateFromToken(claims jwt.MapClaims) error {
	c.UserId = claims["userId"].(string)
	c.ExpiresAt = int64(claims["exp"].(float64))
	return nil
}

func TestGenerateJWT(t *testing.T) {
	jwt, err := GenerateJWT(jwt.MapClaims{
		"userId": "1",
	}, 3)
	if err != nil {
		t.Errorf("There was an error on generating jwt: %v", err)
	}

	if jwt == "" {
		t.Error("There was an error on generating jwt: token is empty")
	}
}

func TestJWTValidation(t *testing.T) {
	jwt, err := GenerateJWT(jwt.MapClaims{
		"userId": "1",
	}, 1)
	if err != nil {
		t.Errorf("There was an error on generating jwt: %v", err)
	}

	claims := TestUserJWTClaims{}
	parsedToken, err := ValidateJWT(jwt, &claims)
	if err != nil {
		t.Errorf("There was an error on validating jwt: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("There was an error on validating jwt: token is not valid")
	}

	if claims.UserId != "1" {
		t.Error("There was an error on validating jwt: userId claim is not correct")
	}

	if claims.ExpiresAt <= time.Now().Unix() {
		t.Error("There was an error on validating jwt: expiration time is not correct")
	}
}
