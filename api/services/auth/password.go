package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func ComparePassword(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}

	return true
}
