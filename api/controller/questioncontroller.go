package controller

import (
	"amirjadhav/go-quiz/api/model"
	"amirjadhav/go-quiz/api/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Startquiz(c *fiber.Ctx) error {
	testQuestions := service.CreateTest()
	return c.JSON(testQuestions)
}

func Submitquiz(c *fiber.Ctx) error {

	userResponse := []model.UserResponse{}

	if err := c.BodyParser(&userResponse); err != nil {
		log.Println("Incorrect user response while submitting test")
	}

	score := service.SubmitTest(userResponse)

	return c.JSON(fiber.Map{"score": score, "percent": (score * 10)})
}
