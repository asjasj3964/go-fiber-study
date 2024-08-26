package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Static("/", "./public") // 정적 파일들이 있는 public 디렉토리에 접근한다.
	app.Listen(":3000")
}
