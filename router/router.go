package router

import (
	"github.com/gofiber/fiber/v3"
	"jwt-in-action/auth"
	"jwt-in-action/middleware"
)

func RegisterRoutes(router fiber.Router) {

	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	loginRouter := router.Group("/login")
	loginRouter.Get("/:userId", func(ctx fiber.Ctx) error {
		token, _ := auth.GenerateJWT(ctx.Params("userId"))
		return ctx.SendString(token)
	})

	userRouter := router.Group("/users")
	userRouter.Use(middleware.CheckLogin())
	userRouter.Get("/get/:username", func(ctx fiber.Ctx) error {
		username := ctx.Params("username")
		return ctx.SendString("Hello, " + username + "!")
	})

}
