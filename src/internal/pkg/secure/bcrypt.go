package secure

import "golang.org/x/crypto/bcrypt"

type BcryptResolver struct{}

func (b *BcryptResolver) Check(h string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(password))
	return err == nil
}
