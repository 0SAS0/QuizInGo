package main

import (
	"QuizInGo/QuizData"
	"fmt"
)

func WelcomeScren() { //Wyswietlanie ekranu witajacego
	fmt.Println("Witaj w Quizie Wiedzy")
	fmt.Println("Spróbuj odpowiedzieć na 7 pytań")
	fmt.Println("Powodzenia !!!")
	fmt.Println("Naciśnij ENTER, aby rozpocząć grę .... ")
	fmt.Scanln()
}
func ShowQuestion(question QuizData.Question) int { // pokazywanie pytan wraz z kategoria i odpowiedziami
	fmt.Printf("\nPytanie za %d pkt.\n", question.Category)
	fmt.Println(question.Content)
	for _, o := range question.Answers {
		fmt.Printf("%d. %s\n", o.Id, o.Content) // wylistowanie petla for odpowiedzi i id odpowiedzi
	}
	fmt.Print("\nWybierz odpowiedź (1-4): ")
	var userAnswer int
	fmt.Scan(&userAnswer) // odpowiedz gracza
	return userAnswer
}
func QuizCompleted() { // pokazywanie ekranu po ukonczeniu quizu
	fmt.Println("\nBrawo, to prawidłowa odpowiedź!")
	fmt.Println("Udało ci się ukończyć cały quiz.")
	fmt.Println("GRATULACJE !!!")
}
func FinishGame() { // Koniec gry
	fmt.Println("\nNiestety, to nie jest prawidłowa odpowiedź.")
	fmt.Println("KONIEC GRY")
}
func CorrectAnswer(category int) { // informacja po podaniu odpowiedzi prawidlowej
	fmt.Println("\nBrawo to poprawna odpowiedź !!!")
	fmt.Printf("Zdobywasz %d pkt.\n", category)
	fmt.Println("Naciśnij ENTER, aby zobaczyć następne pytanie...")
	fmt.Scanf("\n")
}
