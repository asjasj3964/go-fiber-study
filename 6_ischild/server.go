package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	if !fiber.IsChild() { // 현재 프로세스가 Prefork의 결과인지 판단한다.
		fmt.Println("Parent Process")
	} else {
		fmt.Println("Child Process")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with Prefork!")
	})

	app.Listen(":3000")
}
