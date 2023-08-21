package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/josh114/fiberapi/database"
)

func welcome(c *fiber.Ctx) {
	 c.SendString("welcome to my request api")
}

func main (){
	database.connectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":9000"))
} 