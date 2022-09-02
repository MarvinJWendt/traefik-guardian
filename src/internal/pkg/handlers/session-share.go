package handlers

import (
	"time"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func SessionShareRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID := c.Query("id")
		if sessionID == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     auth.SESSION_COOKIE_NAME,
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour),
			SameSite: fiber.CookieSameSiteLaxMode,
			HTTPOnly: true,
		})

		return c.Redirect("/")
	}
}
