package appmodels

import "time"

type TestInput struct {
	Message string `json:"message"`
}

type Answer struct {
	AnswerId  int64  `json:"-"`
	AnswerTxt string `json:"answertxt"`
	IsCorrect bool   `json:"iscorrect"`
}

type Question struct {
	QuestionId int64    `json:"-"`
	Question   string   `json:"question"`
	Answers    []Answer `json:"answers"`
}

type TestOutput struct {
	TestId    int64      `json:"-"`
	Theme     string     `json:"theme"`
	Questions []Question `json:"questions"`
}

type AnswersInput struct {
	TestId       int64 `json:"-"`
	UserId       int64 `json:"-"`
	CorrectCount int64 `json:"correctcount"`
}

type UserScore struct {
	UserId    int64     `json:"userid"`
	TestId    int64     `json:"themeid"`
	Score     int64     `json:"score"`
	CreatedAt time.Time `json:"-"`
}
