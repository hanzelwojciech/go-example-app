package auth

import (
	"github.com/golang-jwt/jwt"
)

type CustomTokenClaims struct {
	*jwt.StandardClaims
	Payload interface{}
}

type TokenCreationResult struct {
	Token     string
	ExpiresAt int64
}
