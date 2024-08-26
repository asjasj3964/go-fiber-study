package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	// immutable 세팅
	app := fiber.New(fiber.Config{
		Immutable: true, // immutable하게 만든다.
	})

	// custom하게 CopyString이라는 함수를 만든다.
	app.Get("/:fiber", func(c *fiber.Ctx) error {
		result := utils.CopyString(c.Params("fiber")) // immutable
		return c.SendString("Copied result: " + result)
	})

	app.Listen(":3000") // 웹 서버를 시작해 포트 3000번에서 요청을 수신한다.
}
