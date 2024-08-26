package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message") // 782 상태 코드와 커스텀 에러 메시지를 반환한다.
	})
	app.Listen(":3000")
}
