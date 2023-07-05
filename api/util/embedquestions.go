package util

import (
	"amirjadhav/go-quiz/api/model"
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed questions.json
var questionsjson []byte

func GetAllQuestions() []model.UserQuestionModel {
	questions := []model.UserQuestionModel{}

	err := json.Unmarshal(questionsjson, &questions)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return nil
	}
	return questions
}

func GetAllQuestionsForCheck() []model.QuestionModel {
	questions := []model.QuestionModel{}

	err := json.Unmarshal(questionsjson, &questions)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return nil
	}
	return questions
}
