package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	// parameter들을 사용할 수 있다.
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})
	app.Listen(":3000")
}
