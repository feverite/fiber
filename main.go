package main

import (
	handler "fiber/handler/helloworld"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		CaseSensitive: true,
		ColorScheme:   fiber.DefaultColors,
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Get("/hello", handler.HelloWorld)
	app.Get("/test", handler.Test)

	micro := fiber.New()
	micro.Get("/hello2", handler.HelloWorld)
	micro.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "hello.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
	app.Mount("/static", micro)

	app.Listen("127.0.0.1:3000")

}
