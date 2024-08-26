package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.BaseURL()) // http://google.com
		fmt.Println(c.Get("X-Custom-Header"))
		return c.SendString("Hello World!")
	})
	// http Request
	req := httptest.NewRequest("GET", "http://google.com", nil)
	req.Header.Set("X-Custom-Header", "hi")
	// http Response
	resp, _ := app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body)) // Hello World!
	}
}
