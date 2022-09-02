package handlers

import (
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CheckRoute(store *session.Store, authDomain string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if !auth.CheckAuthenticated(sess) {
			return c.Redirect("//" + authDomain + "/login?callback=" + c.Hostname())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
