package main

import (
	"qjstest/aaa"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	//
	aaa.AddPeter()

	app.Listen(":3000")
}
