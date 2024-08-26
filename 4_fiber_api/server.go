package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Fiber 애플리케이션 생성 및 설정
	app := fiber.New(fiber.Config{
		Prefork:       true,              // 멀티 코어 환경에서 성능을 향상시킨다.
		CaseSensitive: true,              // 경로에 대한 대소문자 구분을 활성화한다.
		StrictRouting: true,              // 경로에 대한 엄격한 라우팅을 활성화한다. (ex. /name != /name/)
		ServerHeader:  "Fiber",           // 응답 헤더에 "Server: Fiber"를 추가한다.
		AppName:       "Test App v1.0.1", // 애플리케이션 이름을 설정한다.
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to " + app.Config().AppName)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Static("/static", "./public") // 정적 파일들이 있는 public 디렉토리에 접근한다.

	app.Listen(":3000")
}
