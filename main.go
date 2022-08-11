package main

import (
	handler "fiber/handler/helloworld"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
		CaseSensitive: true,
		ColorScheme:   fiber.DefaultColors,
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Get("/", handler.HelloWorld)
	app.Get("/test", handler.Test)

	app.Listen("127.0.0.1:3000")

}
