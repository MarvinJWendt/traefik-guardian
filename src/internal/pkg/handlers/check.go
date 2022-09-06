package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
)

func CheckRoute(store *session.Store, authDomain string) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return fmt.Errorf("could not open store: %w", err)
		}

		if !auth.CheckAuthenticated(sess) {
			return ctx.Redirect("//" + authDomain + "/login?callback=" + ctx.Hostname())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
