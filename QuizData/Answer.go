package QuizData

type Answer struct {
	Base
	IsCorrect bool `json:"CzyPoprawna"`
}
