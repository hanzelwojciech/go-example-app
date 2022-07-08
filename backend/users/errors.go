package users

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

var ErrTokenNotFound = errors.New("token not found")

var ErrEmailExists = errors.New("provided e-mail exists")

var ErrUnauthorized = errors.New("unauthorized")
