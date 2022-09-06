package fiberlog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Call other middlewares first
		err := ctx.Next()
		if err != nil {
			return err
		}

		logrus.Infof(
			"%s | %d | %s %s%s",
			ctx.IP(),
			ctx.Response().StatusCode(),
			ctx.Method(),
			ctx.Hostname(),
			ctx.OriginalURL(),
		)

		return nil
	}
}
