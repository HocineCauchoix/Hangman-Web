package main

import (
	"fmt"
	"hangman-web/hangmanClassic/FileReader"
	"hangman-web/hangmanClassic/HangmanStructure"
	"net/http"
)

const port = ":3000"

func main() {
	hangman.SetRandomWordsFile(FileReader.GetWordsFile())
	HangmanStructure.GetRandomWordFromList(hangman)
	HangmanStructure.ChangeLetter(hangman, HangmanStructure.RevealLetters(hangman))
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/hangman", Hangman)
	http.HandleFunc("/rules", Rules)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/lose", Lose)
	http.HandleFunc("/temp", Temp)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/template/css/"))))
	fmt.Println("http://localhost:3000 - Server started on port :3000")

	http.ListenAndServe(port, nil)

}
