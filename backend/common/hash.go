package common

import "golang.org/x/crypto/bcrypt"

func CreateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

func CompareStringWithHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
