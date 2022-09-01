package handlers

import (
	"github.com/MarvinJWendt/simple-forward-auth/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func LogoutRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		err = auth.Unauthenticate(sess)
		if err != nil {
			return err
		}

		return c.SendString("Logged out")
	}
}
