package hash

import (
	"crypto/sha256"
	"fmt"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

func GenerateHash(value string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(value)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
