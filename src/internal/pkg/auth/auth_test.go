package auth

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/config"
)

func TestGetValidPasswords(t *testing.T) {
	t.Setenv("AUTH_HOST", "auth.example.com")
	t.Setenv("PASSWORDS", "plaintext:pass1|pass2|pass3")

	err := config.Initialize()
	testza.AssertNoError(t, err)

	expectedAlgo := "plaintext"
	expectedPasswords := []string{"pass1", "pass2", "pass3"}

	algorithm, passwords := GetValidPasswords()

	testza.AssertEqual(t, expectedAlgo, algorithm, "algorithm doesn't match")
	testza.AssertEqual(t, expectedPasswords, passwords, "passwords don't match")
}
