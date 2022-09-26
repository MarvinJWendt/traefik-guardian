package auth

import (
	"strings"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const SessionCookieName = "traefik_guardian_session_id"

func GetValidPasswords() []string {
	passwordsRaw := config.Passwords.String()
	return strings.Split(passwordsRaw, "|")
}

func CheckPassword(password string) bool {
	validPasswords := GetValidPasswords()
	for _, validPassword := range validPasswords {
		if password == validPassword {
			return true
		}
	}

	return false
}

func Authenticate(session *session.Session) error {
	session.Set("authenticated", true)
	return session.Save()
}

func Unauthenticate(session *session.Session) error {
	return session.Destroy()
}

func IsAuthenticated(session *session.Session) bool {
	return session.Get("authenticated") != nil
}
