package service

import (
	"amirjadhav/go-quiz/model"
	"amirjadhav/go-quiz/util"
)

var questions = util.GetAllQuestionsForCheck()

func SubmitTest(userAnswers []model.UserResponse) int {
	score := 0
	for _, ur := range userAnswers {
		if checkAnswer(ur.QuestionId, ur.UserAnswer) {
			score++
		}
	}
	return score
}

func checkAnswer(id int, ans string) bool {
	for _, qm := range questions {

		if qm.QuestionId == id {

			if qm.Answer == ans {
				return true
			}
		}
	}
	return false
}
