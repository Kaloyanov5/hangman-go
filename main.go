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
	word = strings.TrimSpace(words[rand.IntN(len(words))])

	output := strings.Split(strings.Repeat("_", len(word)), "")

	for wrong < 6 {
		fmt.Printf("\n%s\n%s\nGuesses: %s\n", hangman[wrong], strings.Join(output, ""), strings.Join(guesses, ", "))
		if strings.Join(output, "") == word {
			fmt.Println("Winner!")
			return
		}

		correct, guess := tryLetter(word)
		if !correct {
			wrong++
		} else {
			for i, char := range word {
				if string(char) == guess {
					output[i] = guess
				}
			}
		}
	}
	fmt.Println(hangman[wrong])
	fmt.Println("Game Over!\nWord was: ", word)
}

func tryLetter(word string) (bool, string) {
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
					if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", rune(s[0])) {
						return errors.New("please enter a lowercase letter a-z")
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

	return strings.Split(string(words), "\n")
}
