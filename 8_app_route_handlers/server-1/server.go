package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.SendString("GET request")
	})
	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("POST request")
	})
	app.Listen(":3000")
}
