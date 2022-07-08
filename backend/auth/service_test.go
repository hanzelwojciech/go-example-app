package auth

import "testing"

func TestCreateToken(t *testing.T) {
	var q struct{}

	token, err := CreateToken(q)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if token == nil {
		t.Error("Token is nil")
	}
}
