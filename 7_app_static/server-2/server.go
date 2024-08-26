package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Download:      false,
		Index:         "seongjin.html",
		CacheDuration: 30 * time.Second,
		MaxAge:        3600,
		ModifyResponse: func(c *fiber.Ctx) error { // 응답을 수정하는 함수
			c.Set("X-Custom-Header", "Fiber static file")
			return nil
		},
		Next: func(c *fiber.Ctx) bool { // 이 미들웨어를 건너 뛸 조건을 설정한다.
			return false
		},
	})

	app.Listen(":3000")
}
