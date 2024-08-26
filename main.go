package main

import (
	"github.com/gofiber/fiber/v3"
	router2 "jwt-in-action/router"
	"log"
)

func main() {

	app := fiber.New()

	router := app.Group("/")

	router2.RegisterRoutes(router)

	log.Fatal(app.Listen(":3000"))

}
