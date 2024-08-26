package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	// TCP 리스너 생성
	ln, err := net.Listen("tcp", ":3000") // 3000 포트에서 들어오는 TCP 연결을 수신한다.
	if err != nil {
		log.Fatalf("failed to listen on port 3000: %v", err)
	}

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key") // 인증서(server.crt)와 개인키(server.key) 로드
	if err != nil {
		log.Fatalf("failed to load certificate: %v", err)
	}

	// TLS 리스너로 변환
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{cer},
	}
	tlsListener := tls.NewListener(ln, tlsConfig) // 기존의 TCP 리스너를 TLS 리스너로 변환한다.

	// Fiber 애플리케이션을 TLS 리스너에 바인딩하고 서버를 시작한다.
	if err := app.Listener(tlsListener); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
