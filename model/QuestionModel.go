package model

type UserQuestionModel struct {
	QuestionId int    `json:"questionId"`
	Question   string `json:"question"`
	Options    `json:"options"`
	Answer     string `json:"-"`
}

type Options struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
}

type QuestionModel struct {
	QuestionId int    `json:"questionId"`
	Question   string `json:"question"`
	Options    `json:"options"`
	Answer     string `json:"answer"`
}
