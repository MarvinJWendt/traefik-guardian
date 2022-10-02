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
	hasher := sha512.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
