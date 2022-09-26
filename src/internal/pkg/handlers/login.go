package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/config"
)

func LoginAPI(store *session.Store) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		password := ctx.FormValue("password")

		if !auth.CheckPassword(password) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		sess, err := store.Get(ctx)
		if err != nil {
			return err
		}

		err = auth.Authenticate(sess)
		if err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func LoginRoute(store *session.Store) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		callback := ctx.Query("callback")

		sess, err := store.Get(ctx)
		if err != nil {
			return err
		}

		if auth.IsAuthenticated(sess) {
			return ctx.Redirect("//" + callback + "/traefik-guardian-session-share?id=" + sess.ID())
		}

		return ctx.Render("login", fiber.Map{
			"Callback":  callback,
			"SessionID": sess.ID(),
			"Title":     config.LoginPageTitle.Value,
		})
	}
}
