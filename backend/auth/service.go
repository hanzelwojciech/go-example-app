package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwtKey      []byte
	initialized bool = false
)

func initialize() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
	initialized = true
}

func CreateToken(payload interface{}) (*TokenCreationResult, error) {
	if !initialized {
		initialize()
	}

	method := jwt.New(jwt.GetSigningMethod("HS256"))
	expiresAt := time.Now().Add(time.Minute * 5).Unix()

	method.Claims = &CustomTokenClaims{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		&payload,
	}

	token, err := method.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}

	return &TokenCreationResult{Token: token, ExpiresAt: expiresAt}, nil
}

func VerifyDecodeToken(token string) (*CustomTokenClaims, error) {
	if !initialized {
		initialize()
	}

	data, err := jwt.ParseWithClaims(token, &CustomTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		fmt.Println("test", err.Error())

		return nil, err
	}

	return data.Claims.(*CustomTokenClaims), nil
}
