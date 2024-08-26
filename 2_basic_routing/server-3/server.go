package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	// wildcard를 넣을 수 있다.
	app.Get("/api/*", func(c *fiber.Ctx) error { // /api/ 하위의 모든 경로를 하나의 handler에서 처리한다.
		return c.SendString("API path: " + c.Params("*"))
	})
	app.Listen(":3000")
}
