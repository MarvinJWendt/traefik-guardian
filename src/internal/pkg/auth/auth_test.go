package auth

import (
	"os"
	"reflect"
	"testing"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/config"
)

func TestGetValidPasswords(t *testing.T) {
	os.Setenv("AUTH_HOST", "*")
	os.Setenv("PASSWORDS", "plaintext:pass1|pass2|pass3")
	if err := config.Initialize(); err != nil {
		t.Fatalf("couldn't initialize config: %s", err)
	}

	expectedAlgo := "plaintext"
	expectedPasswords := []string{"pass1", "pass2", "pass3"}

	algorithm, passwords := GetValidPasswords()
	if expectedAlgo != algorithm {
		t.Fatalf("couldn't get algorithn:\ngot: %v\nwanted: %v\n", algorithm, expectedAlgo)
	}
	if !reflect.DeepEqual(expectedPasswords, passwords) {
		t.Fatalf("couldn't get valid passwords:\ngot: %v\nwanted: %v\n", passwords, expectedPasswords)
	}
}
