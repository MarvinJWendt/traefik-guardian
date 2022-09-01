package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

func SessionShareRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID := c.Query("id")
		if sessionID == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "simple_forward_auth_session_id",
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour),
			SameSite: "lax",
		})

		return c.Redirect("/")
	}
}
