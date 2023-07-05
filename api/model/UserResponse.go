package model

type UserResponse struct {
	QuestionId int    `json:"questionId"`
	UserAnswer string `json:"userAnswer"`
}
