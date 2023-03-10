package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashSHA256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
