package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"jacc/config"
)

func main() {
	config, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	app := startServer(config)
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(config config.Config) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, %s!", config.Name))
	})

	return app
}
