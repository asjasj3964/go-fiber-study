package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/random"
)

func main() {
	app := fiber.New()
	// 1
	app.Use(func(c *fiber.Ctx) error { // 모든 요청에 대해 실행되는 기본 미들웨어 설정
		return c.Next()
	})

	// 2
	app.Use("/api", func(c *fiber.Ctx) error { // 특정 경로에 미들웨어 적용 (/api로 시작하는 모든 요청에 대해 미들웨어 적용)
		return c.Next()
	})

	// 3
	app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error { // 다중 경로에 미들웨어 적용
		return c.Next()
	})

	// 4
	app.Use("/api", func(c *fiber.Ctx) error { // 여러 handler를 가진 미들웨어
		c.Set("X-Custom-Header", random.String(32)) // 헤더 설정
		return c.Next()
	}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Listen(":3000")
}
