# Hangman Game in Go

A simple command-line implementation of the classic Hangman game using Go and `huh` TUI package.

## Description

This Hangman game features:
- ASCII art visualization of the hangman
- Word selection from a text file
- Interactive letter input with validation
- Progress tracking and game state display

## Prerequisites

- Go 1.22 or higher
- `github.com/charmbracelet/huh` package

## Installation

```bash
git clone https://github.com/Kaloyanov5/hangman-go.git
cd hangman-go
go mod tidy
```

## Usage
1. Ensure you have a words.txt file with one word per line
2. Run the game:

```bash
go run main.go
```

## Game Rules
- You have 6 attempts to guess the word
- Enter one letter at a time
- Already guessed letters cannot be used again
- The game ends when you either:
  - Correctly guess the word (win)
  - Run out of attempts (lose)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
