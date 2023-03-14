package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func HashSHA256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

func HashSHA512(input string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
}
