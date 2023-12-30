package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := startServer()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Potential Employer 👋!")
	})

	return app
}
