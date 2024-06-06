package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing the password: %v", err)
	}

	if hash == "" {
		t.Errorf("the hash is empty")
	}

	if hash == password {
		t.Errorf("the password is not hashed")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing the password: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Errorf("error while validating the password hash")
	}

	if CheckPasswordHash("password2", hash) {
		t.Errorf("validate the wrong password")
	}
}
