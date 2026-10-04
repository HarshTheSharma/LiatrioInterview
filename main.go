package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getMessage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "My name is Harsh Sharma",
		"timestamp": time.Now().UnixMilli(),
	})
}

func main() {
	app := fiber.New()
	app.Get("/", getMessage)
	log.Fatal(app.Listen(":8080"))
}