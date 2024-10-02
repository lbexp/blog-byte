package bcrypt_utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(value string) (string, error) {
	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), cost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashAndValue(hash string, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
