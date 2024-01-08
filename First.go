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
			ValiuEasy(H)

			Forme(H)
		} else if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
			LoadGame("save.txt", H)
			Forme(H)
		} else {
			os.Exit(1)
		}
	} else {
		ValiuEasy(H)
		Forme(H)
	}
}

func ValiuEasy(H *HangManData) {
	H.File = "words/words.txt"
	H.ToFind = Word(H)
	H.Word = UnderscoreWord(H)
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	H.Attempts = 10
	H.Num = 0

}

func ValiuNormal(H *HangManData) {
	H.File = "words/words2.txt"
	H.ToFind = Word(H)
	H.Word = UnderscoreWord(H)
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	H.Attempts = 10
	H.Num = 0

}

func ValiuHard(H *HangManData) {
	H.File = "words/words3.txt"
	H.ToFind = Word(H)
	H.Word = UnderscoreWord(H)
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	H.Attempts = 10
	H.Num = 0

}
