package Hangman

import (
	"bufio"
	"log"
	"os"
)

// function used when the given letter is not in the word to be guessed
func False(H *HangManData) string {
	Drow := ""
	file, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	Line := 0
	for fileScanner.Scan() {
		if H.Attempts >= 10 {
			if Line >= 72 && Line <= 80 {
				Drow += fileScanner.Text() + "\n"
			}
		} else if Line >= H.HangmanPositions[9-H.Attempts+1] && Line < H.HangmanPositions[9-H.Attempts] && H.Attempts <= 9 {
			Drow += fileScanner.Text() + "\n"
		}
		Line++
	}
	H.Drow = Drow
	return H.Drow
}
