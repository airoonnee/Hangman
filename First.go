package Hangman

import (
	"flag"
	"fmt"
	"os"
)

func Test() string {
	fmt.Print("ca marche bg")
	ok := "bien jouer bg"

	return ok
}

func First() {
	var H *HangManData //pointer initialisation
	H = new(HangManData)

	flag.String("startWith", "default", "File name to start with")
	flag.String("letterFile", "default", "File name to choose ASCII")
	flag.Parse()
	// checks whether the player wants to start the game with ASCII Art
	// or whether he wants to start the game he has saved
	if len(os.Args) == 3 {
		if os.Args[2] == "--letterFile" {
			Valiu(H)

			Forme(H)
		} else if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
			LoadGame("save.txt", H)
			Forme(H)
		} else {
			os.Exit(1)
		}
	} else {
		Valiu(H)
		Forme(H)
	}
}

func Valiu(H *HangManData) {
	H.ToFind = Word()
	H.Word = UnderscoreWord(H)
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	H.Attempts = 0
}
