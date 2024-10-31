package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("There was an error on hashing the password: %v", err)
	}

	if hash == "" {
		t.Error("There was an error on hashing the password: hash is empty")
	}

	if hash == "password" {
		t.Error(
			"There was an error on hashing the password: hash is equal to the original password",
		)
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("There was an error on hashing the password: %v", err)
	}

	if !ComparePassword("password", hash) {
		t.Error(
			"There was an error on comparing the password: expected hash to be matched with the given password 'password'",
		)
	}

	if ComparePassword("password123", hash) {
		t.Error(
			"There was an error on comparing the password: expected hash to not be matched with the given password 'password123'",
		)
	}
}
