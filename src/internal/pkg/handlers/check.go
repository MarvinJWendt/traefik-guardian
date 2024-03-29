package handlers

import (
	"fmt"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/auth"
)

func CheckRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return fmt.Errorf("could not open store: %w", err)
		}

		if ctx.GetReqHeaders()["Guardian-Password"] != "" {
			if !auth.CheckPassword(ctx.GetReqHeaders()["Guardian-Password"]) {
				return ctx.SendStatus(fiber.StatusUnauthorized)
			}

			return ctx.SendStatus(fiber.StatusOK)
		}

		if !auth.IsAuthenticated(sess) {
			return ctx.Redirect("//" + config.AuthHost.String() + "/login?callback=" + ctx.Hostname())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
