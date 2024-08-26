package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:fiber", handler) // 동적 경로 :fiber 설정
	app.Listen(":3000")
}

func handler(c *fiber.Ctx) error {
	// fiber.Ctx로부터 반환된 값 = not immutable (default)
	// -> 일부 값들이 requests 사이에 재사용될 수 있다.
	// 오직 handler 안에서만 context value들을 사용해야 하고,
	// 어떠한 reference들을 유지하지 말아야 한다.
	// handler에서 return을 하면 context로부터 얻은 어떠한 값들도 미래의 request에 다시 재사용 될 것이고
	// 사용자의 제어에 벗어나 변경될 것이다.
	result := c.Params("fiber")

	// 해당 값들을 handler 밖에서도 보존하고 싶다면
	// copy를 사용하여 underlying buffer의 복사본을 만들어야 한다.
	buffer := make([]byte, len(result)) // result 문자열 길이 만큼의 byte slice를 만든다.
	copy(buffer, result)                // 복사본을 만든다.
	resultCopy := string(buffer)
	// string엔 buffer가 있고 이를 변경하면 immutable한 string도 변경된다.

	return c.SendString("Copied result: " + resultCopy)
}
