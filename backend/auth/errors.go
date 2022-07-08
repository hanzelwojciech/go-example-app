package auth

import "errors"

var ErrTokenExpired = errors.New("token expired")

var ErrUnauthorized = errors.New("unauthorized")
