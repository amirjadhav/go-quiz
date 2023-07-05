package service

import (
	"amirjadhav/go-quiz/api/model"
	"amirjadhav/go-quiz/api/util"
	"math/rand"
)

func CreateTest() []model.UserQuestionModel {

	questions := util.GetAllQuestions()
	// create a set of 10 random numbers
	var randomNumberSet []int

	for {
		if len(randomNumberSet) == 10 {
			break
		}
		num := rand.Intn(100) + 1
		if sliceContains(randomNumberSet, num) {
			continue
		}

		randomNumberSet = append(randomNumberSet, num)
	}

	// select respective question and make new teststruct
	// which have question with id and without answer

	testQuestions := []model.UserQuestionModel{}

	for _, questionNumber := range randomNumberSet {

		// get respective struct
		testQuestions = append(testQuestions, getQuestion(questions, questionNumber))
	}
	return testQuestions
}

func sliceContains(numbers []int, num int) bool {
	for _, v := range numbers {
		if v == num {
			return true
		}
	}
	return false
}

func getQuestion(questions []model.UserQuestionModel, questionNumber int) model.UserQuestionModel {
	for _, question := range questions {
		if question.QuestionId == questionNumber {
			return question
		}
	}
	return model.UserQuestionModel{}
}
