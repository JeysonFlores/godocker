package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Mazatlán",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/json", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&Response{ID: 1, Name: "xd", Msg: "asas"})
	})

	app.Get("/error", func(ctx *fiber.Ctx) error {

		return ctx.Status(404).JSON(&fiber.Map{
			"error": true,
			"msg":   "this route always throws error",
		})
	})

	app.Listen(":8080")
}
