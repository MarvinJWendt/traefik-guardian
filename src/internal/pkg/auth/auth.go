package auth

import (
	"strings"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const SessionCookieName = "traefik_guardian_session_id"

func GetValidPasswords() (string, []string) {
	// Schema: "algorithm:pass1|pass2|pass3"
	passwordsRaw := config.Passwords.String()
	algorithm := strings.Split(passwordsRaw, ":")[0]
	passwords := strings.Split(strings.Split(passwordsRaw, ":")[1], "|")
	return algorithm, passwords
}

func CheckPassword(password string) bool {
	_, validPasswords := GetValidPasswords()
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
