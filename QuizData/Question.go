package QuizData

type Question struct {
	Base
	Category int      `json:"Kategoria"`
	Answers  []Answer `json:"Odpowiedzi"`
}
