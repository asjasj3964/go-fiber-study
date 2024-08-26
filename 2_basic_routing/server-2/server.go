package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	// optional parameter를 넣을 수 있다.
	app.Get("/:name", func(c *fiber.Ctx) error {
		if c.Params("name") != "SeongJin" {
			return c.SendString("Where is SeongJin?")
		}
		return c.SendString("Hello " + c.Params("name"))
	})
	app.Listen(":3000")
}
