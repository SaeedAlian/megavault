package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"megavault/api/config"
	types_user "megavault/api/types"
	"megavault/api/utils"
)

type JWTClaims interface {
	PopulateFromToken(claims jwt.MapClaims) error
}

func WithJWTAuth(handler http.HandlerFunc, store types_user.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		claims, token, err := ValidateJWT[*types_user.UserJWTClaims](tokenStr)
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

func GenerateJWT(claims jwt.MapClaims, expiresAtInMinutes int) (string, error) {
	expiration := time.Second * time.Duration(expiresAtInMinutes*60)

	tokenClaims := jwt.MapClaims{}

	for k, v := range claims {
		tokenClaims[k] = v
	}

	tokenClaims["expiresAt"] = time.Now().Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	secret := []byte(config.Env.JWTSecret)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT[T JWTClaims](tokenString string) (T, *jwt.Token, error) {
	claims := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)

	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil {
		return claims, nil, err
	}

	mapClaims := parsed.Claims.(jwt.MapClaims)

	if err := claims.PopulateFromToken(mapClaims); err != nil {
		return claims, nil, err
	}

	return claims, parsed, nil
}
