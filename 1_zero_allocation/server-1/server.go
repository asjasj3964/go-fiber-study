package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()                      // fiber 애플리케이션 인스턴스 생성
	app.Get("/", func(c *fiber.Ctx) error { // 루트 경로 '/''에 대한 GET 요청을 처리하는 handler
		return c.SendString("Hello World!") // 클라이언트에게 응답을 보낸다.
	})
	// func(c * fiber.Ctx) error { ... } : 요청을 처리하는 handler 함수
	// c * fiber.Ctx : 요청 및 응답 컨텍스트를 나타내는 매개변수

	app.Listen(":3000") // 웹 서버를 시작해 포트 3000번에서 요청을 수신한다.
}
