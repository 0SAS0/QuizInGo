package main

func main() {
	backend := NewGameService()
	WelcomeScren()
	for {
		backend.RandSelQuestion()
		userAnswer := ShowQuestion(backend.ActualQuestion)
		if backend.CheckUserAnswers(userAnswer) {
			if backend.IsLastQuestion() {
				QuizCompleted()
				break
			} else {
				CorrectAnswer(backend.ActualCategory)
				backend.LiftCategory()
			}
		} else {
			FinishGame()
			break
		}
	}
}
