package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	id := 1
	token, err := CreateJWT(id)
	if err != nil {
		t.Errorf("error while creating the token: %v", err)
	}

	if len(token) < 100 {
		t.Errorf("token too short")
	}
}

func TestValidateJWT(t *testing.T) {
	id := 1
	tokenStr, err := CreateJWT(id)
	if err != nil {
		t.Errorf("error while creating the token: %v", err)
	}

	token, err := validateJWT(tokenStr)
	if err != nil {
		t.Errorf("failed to validate the token: %v", err)
	}

	if !token.Valid {
		t.Errorf("token is invalid")
	}
}
