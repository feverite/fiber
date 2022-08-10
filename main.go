package main

import (
	handler "fiber/handler/helloworld"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.HelloWorld)

	app.Listen("127.0.0.1:3000")

}
