package middleware

import (
	"github.com/gofiber/fiber/v3"
	"jwt-in-action/auth"
)

func CheckLogin() fiber.Handler {

	return func(ctx fiber.Ctx) error {

		headerMap := ctx.GetReqHeaders()

		authHead := headerMap["Authorization"]

		if authHead == nil || len(authHead) == 0 {
			return ctx.Status(401).SendString("Unauthorized")
		}

		token := authHead[0]

		if token == "" {
			return ctx.Status(401).SendString("Unauthorized")
		}

		if _, err := auth.VerifyJWT(token); err != nil {
			return ctx.Status(401).SendString("Unauthorized")
		}

		return ctx.Next()

	}

}
