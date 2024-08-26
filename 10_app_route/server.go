package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Route("/test", func(api fiber.Router) {
		api.Get("/seongjin", handler).Name("seongjin") // /test/seongjin, 경로에 이름을 붙인다. -> test.seongjin
		api.Get("/heejin", handler).Name("heejin")     // /test/heejin
	}, "test.") // 라우트 그룹 이름을 지정한다.

	log.Fatal(app.Listen(":3000"))
}

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello " + c.Path())
}
