package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/MarvinJWendt/traefik-guardian/src/internal/pkg/auth"
)

func SessionShareRoute() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sessionID := ctx.Query("id")
		if sessionID == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Cookie(&fiber.Cookie{
			Name:     auth.SessionCookieName,
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour),
			SameSite: fiber.CookieSameSiteLaxMode,
			HTTPOnly: true,
		})

		return ctx.Redirect("/")
	}
}
