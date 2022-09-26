package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/auth"
)

func LogoutRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return err
		}

		err = auth.Unauthenticate(sess)
		if err != nil {
			return err
		}

		return ctx.SendString("Logged out")
	}
}
