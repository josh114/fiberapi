package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func welcome(c *fiber.Ctx) {
	 c.SendString("welcome to my request api")
}

func main (){
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":9000"))
}