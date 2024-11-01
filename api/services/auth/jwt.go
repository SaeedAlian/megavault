package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"megavault/api/config"
	"megavault/api/types/user"
	"megavault/api/utils"
)

type JWTClaims interface {
	PopulateFromToken(claims jwt.MapClaims) error
}

func WithJWTAuth(handler http.HandlerFunc, store types_user.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		claims := types_user.UserJWTClaims{}
		token, err := ValidateJWT(tokenStr, &claims)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			utils.WriteErrorInResponse(w, http.StatusUnauthorized, "Failed to validate token")
			return
		}

		if !token.Valid {
			log.Printf("invalid token received")
			utils.WriteErrorInResponse(w, http.StatusUnauthorized, "Invalid token received")
			return
		}

		userId := claims.UserId

		u, err := store.GetUserById(userId)
		if u == nil || err != nil {
			log.Printf("invalid token received")
			utils.WriteErrorInResponse(w, http.StatusUnauthorized, "Invalid token received")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", u.Id)
		r = r.WithContext(ctx)

		handler(w, r)
	}
}

func GenerateJWT(claims jwt.MapClaims, expiresAtInMinutes float64) (string, error) {
	expiration := time.Minute * time.Duration(expiresAtInMinutes)

	tokenClaims := jwt.MapClaims{}

	for k, v := range claims {
		tokenClaims[k] = v
	}

	tokenClaims["exp"] = time.Now().UTC().Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	secret := []byte(config.Env.JWTSecret)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT[T JWTClaims](tokenString string, claims T) (*jwt.Token, error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	mapClaims := parsed.Claims.(jwt.MapClaims)

	if err := claims.PopulateFromToken(mapClaims); err != nil {
		return nil, err
	}

	return parsed, nil
}
