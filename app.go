package main

import (
	"os"
	"time"

	"github.com/MarvinJWendt/simple-forward-auth/internal/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html"
)

func main() {
	authDomain := os.Getenv("AUTH_DOMAIN")
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	store := session.New(session.Config{
		Expiration:   24 * time.Hour,
		KeyLookup:    "cookie:simple_forward_auth_session_id",
		CookiePath:   "/",
		KeyGenerator: utils.UUID,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		authenticated := sess.Get("authenticated")
		if authenticated == nil {
			return c.SendString("Not authenticated")
		}

		return c.SendString("Authenticated")
	})

	app.Get("/simple-forward-auth-session-share", func(c *fiber.Ctx) error {
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
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		callback := c.Query("callback")

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		return c.Render("login", fiber.Map{
			"Callback":  callback,
			"SessionID": sess.ID(),
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
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

	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		err = auth.Unauthenticate(sess)
		if err != nil {
			return err
		}

		return c.SendString("Logged out")
	})

	app.Get("/check", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		authenticated := sess.Get("authenticated")
		if authenticated == nil {
			return c.Redirect("//" + authDomain + "/login?callback=" + c.Hostname())
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/session/get", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		return c.SendString(sess.ID())
	})

	app.Listen(":80")
}
