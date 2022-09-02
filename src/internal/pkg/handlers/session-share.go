package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func SessionShareRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID := c.Query("id")
		if sessionID == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "simple_forward_auth_session_id",
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour),
			SameSite: fiber.CookieSameSiteLaxMode,
			HTTPOnly: true,
		})

		return c.Redirect("/")
	}
}
