package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	// 하나의 경로에 여러 개의 디렉토리를 제공한다.
	app.Static("/", "./files")
	app.Static("/", "./public")

	app.Static("/static", "./public") // 가상 path prefix 설정 (= 실제 filesystem엔 없지만 route 과정에서의 path를 설정해놓은 것)
	app.Listen(":3000")
}
