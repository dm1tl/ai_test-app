package appmodels

import "time"

type TestInput struct {
	Message string `json:"message"`
}

type Answer struct {
	AnswerTxt string `json:"answertxt"`
	IsCorrect bool   `json:"iscorrect"`
}

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
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
	ThemeId   int64     `json:"themeid"`
	Score     int64     `json:"score"`
	CreatedAt time.Time `json:"-"`
}

type Themes struct {
	ThemeId int64  `json:"themeid"`
	Name    string `json:"name"`
}
