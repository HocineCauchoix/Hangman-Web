package main

import (
	"fmt"
	"hangman-web/hangmanClassic/HangmanStructure"
	"hangman-web/hangmanClassic/UserInput"
	"net/http"
	"strconv"
	"text/template"
)

var (
	newString string
	attempts  int
	isCorrect bool = true
	beginGame bool = true
	endGame   bool = false
	won       bool = false
)

type indexPageData struct {
	WordToFind string
	Attempts   int
	ImageName  string
	IsCorrect  bool
	BeginGame  bool
	EndGame    bool
	Won        bool
}

var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

var hangman = new(HangmanStructure.HangmanData)

func Hangman(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("html/template/hangman.html"))
	newString = ""
	for _, letter := range hangman.GetWord() {
		newString += letter
	}
	if newString == hangman.GetWordToFind() {
		endGame = true
		won = true
	}

	data := indexPageData{
		WordToFind: newString,
		Attempts:   10 - attempts,
		ImageName:  strconv.Itoa(attempts),
		IsCorrect:  isCorrect,
		BeginGame:  beginGame,
		EndGame:    endGame,
		Won:        won,
	}
	template.Execute(w, data)
}
func Win(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "win")

}

func Lose(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "lose")

}

func Rules(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "rules")
}

func Temp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	letter := r.FormValue("letter")

	if !UserInput.IsLetterCorrect(letter, hangman) {
		attempts += 1
		isCorrect = false
		beginGame = false
		if attempts == 10 {
			endGame = true
		}
	} else {
		isCorrect = true
		beginGame = false
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
