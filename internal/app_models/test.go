package appmodels

type TestInput struct {
	Message string `json:"message"`
}

type TestOutput struct {
	Theme     string      `json:"theme"`
	Questions []Questions `json:"questions"`
}

type Questions struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type AnswersInput struct {
	Count int64 `json:"count"`
}
