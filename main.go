package main

import (
	"amirjadhav/go-quiz/controller"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", controller.HelloWorld)

	app.Get("/api/startquiz", controller.Startquiz)

	app.Post("/api/submitquiz", controller.Submitquiz)

	log.Fatal(app.Listen(":3000"))
}
