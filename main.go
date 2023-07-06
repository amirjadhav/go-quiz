package main

import (
	"amirjadhav/go-quiz/controller"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", controller.HelloWorld)

	app.Get("/api/startquiz", controller.Startquiz)

	app.Post("/api/submitquiz", controller.Submitquiz)

	log.Fatal(app.Listen(":3000"))
}
