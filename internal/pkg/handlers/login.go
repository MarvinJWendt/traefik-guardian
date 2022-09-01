package handlers

import (
	"github.com/MarvinJWendt/simple-forward-auth/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func LoginAPI(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		password := c.FormValue("password")

		if !auth.CheckPassword(password) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		err = auth.Authenticate(sess)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func LoginRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		callback := c.Query("callback")

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		return c.Render("login", fiber.Map{
			"Callback":  callback,
			"SessionID": sess.ID(),
		})
	}
}
