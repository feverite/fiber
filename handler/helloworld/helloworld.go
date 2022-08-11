package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}
	return c.SendString("Hello, World!")
}

func Test(c *fiber.Ctx) error {

	for i := 0; i < 100000; i++ {
		go func(i int) {
			fmt.Println(i)
			a := fiber.AcquireAgent()
			req := a.Request()
			req.Header.SetMethod(fiber.MethodGet)
			req.SetRequestURI("http://localhost:3000/")
			if err := a.Parse(); err != nil {
				panic(err)
			}

			a.Bytes()
		}(i)

		// fmt.Println(code)
		// fmt.Println(string(body))
	}

	return c.SendStatus(200)
}
