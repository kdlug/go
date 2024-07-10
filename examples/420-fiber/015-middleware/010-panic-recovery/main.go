package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Get("/", helloworld)
	app.Listen(":3002")
}

func helloworld(c *fiber.Ctx) error {
	panic("Uuuups!")
	return c.SendString("Hello world")
}
