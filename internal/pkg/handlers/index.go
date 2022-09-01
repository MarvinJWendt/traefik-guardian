package handlers

import (
	"github.com/MarvinJWendt/simple-forward-auth/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func IndexRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if !auth.CheckAuthenticated(sess) {
			return c.SendString("Not authenticated")
		}

		return c.SendString("Authenticated")
	}
}
