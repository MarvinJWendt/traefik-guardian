package handlers

import (
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/config"
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

		if auth.CheckAuthenticated(sess) {
			return c.Redirect("//" + callback + "/traefik-auth-provider-session-share?id=" + sess.ID())
		}

		return c.Render("login", fiber.Map{
			"Callback":  callback,
			"SessionID": sess.ID(),
			"Title":     config.LOGIN_PAGE_TITLE.Value,
		})
	}
}
