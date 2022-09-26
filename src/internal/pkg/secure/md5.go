package secure

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Resolver struct{}

func (m *MD5Resolver) Check(h string, password string) bool {
	return h == GetMD5Hash(password)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
