package api

import (
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Authenticated(store *session.Store) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return err
		}
		if !auth.CheckAuthenticated(sess) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
