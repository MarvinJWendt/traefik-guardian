package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
)

func IndexRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return err
		}

		if !auth.IsAuthenticated(sess) {
			return ctx.SendString("Not authenticated")
		}

		return ctx.SendString("Authenticated")
	}
}
