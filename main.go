package main

// ////////////////////////////////////////////////////////////////
// Grupa: 1 Łódź                                                 //
// Wykonał: Szymon Niewiadomski 122861 | Daniel Stasiak 122896   //
// ////////////////////////////////////////////////////////////////
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
