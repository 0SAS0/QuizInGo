package main

import (
	"QuizInGo/QuizData"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"time"
)

type GameService struct {
	Random         *rand.Rand
	listQuestion   []QuizData.Question
	ActualCategory int
	Category       []int
	ActualQuestion QuizData.Question
	ActualIndex    int
}

func NewGameService() *GameService { // funkcja tworząca nową gre
	service := &GameService{
		Random: rand.New(rand.NewSource(time.Now().UnixNano())), // generator liczb losowych
	}
	service.CreateListQuestion()    // wywolanie funkcji wczytywania listy pytań
	service.DownloadCategory()      // Pobranie kategori
	if len(service.Category) == 0 { // zwykly if do sprawdzenia czy istnieja kategorie
		fmt.Println("You have no category!")
	}
	service.ActualCategory = service.Category[service.ActualIndex] // ustawienie kategorii na 100
	return service
}
func (service *GameService) RandSelQuestion() { // funkcja losowania pytania z akutalnej kategorii
	var QuestionFromActuallyCategory []QuizData.Question
	for _, p := range service.listQuestion { // przejscie przez wsystkie pytania oraz sprawdzenie czy pytanie nalezy do odpowiedniej kategorii i dodanie go
		if p.Category == service.ActualCategory {
			QuestionFromActuallyCategory = append(QuestionFromActuallyCategory, p)
		}
	}
	if len(QuestionFromActuallyCategory) == 0 {
		fmt.Println("You have no category!")
	}
	Index := service.Random.Intn(len(QuestionFromActuallyCategory)) // losowanie indexu pytania
	DrawQuestion := QuestionFromActuallyCategory[Index]
	rand.Shuffle(len(DrawQuestion.Answers), func(i, j int) { // mieszanie kolejnosci odpowiedzi aby sie nie powtarzaly
		DrawQuestion.Answers[i], DrawQuestion.Answers[j] = DrawQuestion.Answers[j], DrawQuestion.Answers[i]
	})
	for i := range DrawQuestion.Answers { // numeracja pętlą for odpowiedzi
		DrawQuestion.Answers[i].Id = i + 1
	}
	service.ActualQuestion = DrawQuestion
}
func (service *GameService) CheckUserAnswers(userAnswer int) bool { // funckja sprawdzajaca czy odpowiedz gracza jest poprawna
	for _, q := range service.ActualQuestion.Answers {
		if q.Id == userAnswer {
			return q.IsCorrect
		}
	}
	return false
}
func (service *GameService) IsLastQuestion() bool { // funkcja sprawdzjaca czy aktualne pytanie jest ostatnie
	return service.ActualIndex == len(service.Category)-1
}
func (service *GameService) LiftCategory() { // przechodzi do wyzszej kategorii o ile jest dostepna
	if service.ActualIndex+1 < len(service.Category) {
		service.ActualIndex++
		service.ActualCategory = service.Category[service.ActualIndex]
	} else {
		fmt.Println("Out of scope") //  na wypadek wyjscia z zakresu
	}

}
func (service *GameService) CreateListQuestion() { // wczytuje liste pytan z pliku questions_pl.json
	path := "questions_pl.json"
	text, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(text, &service.listQuestion); err != nil {
		panic(err)
	}
}
func (service *GameService) DownloadCategory() { // Pobiera dostepne kategorie i sortuje je od najmniejszego do najwiekszego
	CategoryMap := make(map[int]bool)        // mapa do przechowywania kategorii
	for _, q := range service.listQuestion { // przejscie przez liste pytan i dla kazdego pytania ze dana kategoria istnieje przypisujać wartoość true
		CategoryMap[q.Category] = true
	}

	for k := range CategoryMap { // przechodzenie przez klucze w mapie i dodanie kategorii do listy kategorii
		service.Category = append(service.Category, k)
	}
	sort.Ints(service.Category) // funkcja sortujaca
	if len(service.Category) == 0 {
		fmt.Println("Category Not Found")
	}
}
