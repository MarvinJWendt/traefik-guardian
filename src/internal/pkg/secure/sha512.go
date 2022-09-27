package secure

import (
	"crypto/sha512"
	"encoding/hex"
)

type SHA512Resolver struct{}

func (s *SHA512Resolver) Check(h string, password string) bool {
	return h == GetSHA512Hash(password)
}

func GetSHA512Hash(text string) string {
	hash := sha512.New().Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
