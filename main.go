package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

var (
	hangman = []string{"+---+\n|   |\n|\n|\n|\n|\n=========", "+---+\n|   |\n|   O\n|\n|\n|\n=========",
		"+---+\n|   |\n|   O\n|   |\n|\n|\n=========", "+---+\n|   |\n|   O\n|  /|\n|\n|\n=========",
		"+---+\n|   |\n|   O\n|  /|\\\n|\n|\n=========", "+---+\n|   |\n|   O\n|  /|\\\n|  /\n|\n=========",
		"+---+\n|   |\n|   O\n|  /|\\\n|  / \\\n|\n========="}
	word     string
	filename = "words.txt"
	wrong    = 0
	guesses  []string
)

func main() {
	words := getWords(filename)
	word = words[rand.IntN(len(words))]

	output := strings.Split(strings.Repeat("_", len(word)), "")

	for wrong <= 6 {
		fmt.Println(hangman[wrong])
		fmt.Println(strings.Join(output, ""))
		if strings.Join(output, "") == word {
			fmt.Println("winner")
			return
		}
		correct, guess := startGame(word)
		if !correct {
			fmt.Println("Wrong guess!")
			wrong++
		} else {
			output[strings.Index(word, guess)] = guess
		}
	}
}

func startGame(word string) (bool, string) {
	guess := ""
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter a letter").
				Value(&guess).
				Validate(func(s string) error {
					if len(s) != 1 {
						return errors.New("please enter a single letter")
					}
					if slices.Contains(guesses, guess) {
						return errors.New("you have already entered this letter")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	guesses = append(guesses, guess)

	return strings.Contains(word, guess), guess
}

func getWords(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	words, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	return strings.Split(string(words), ",")
}
