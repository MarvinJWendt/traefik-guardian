package auth

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetValidPassword() string {
	return os.Getenv("PASSWORD")
}

func CheckPassword(password string) bool {
	return password == GetValidPassword()
}

func Authenticate(session *session.Session) error {
	session.Set("authenticated", true)
	return session.Save()
}

func Unauthenticate(session *session.Session) error {
	return session.Destroy()
}

func CheckAuthenticated(session *session.Session) bool {
	return session.Get("authenticated") != nil
}
