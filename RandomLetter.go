package Hangman

import (
	"math/rand"
	"strings"
	"time"
)

// transform the word to guess by hiding letters with _.
func UnderscoreWord(H *HangManData) string {
	Word := H.ToFind
	// Calculate how many letters reveal
	revealCount := len(Word)/2 - 1
	revealed := make([]string, len(Word))
	underscoreWord := []string{}
	for k := 0; k < len(Word); k++ { // create an array of strings the length of the word, where each cell is a _.
		underscoreWord = append(underscoreWord, "_")
	}
	for p := 0; p < len(Word); p++ { //create a string table of the word find, with each letter in a box
		revealed = strings.Split(Word, "")
	}
	H.Letter = ""
	var Wordend string
	TabLetter := []string{}
	for j := 0; j < revealCount; j++ {
		rand.Seed(time.Now().UnixNano())
		letter := rand.Intn(len(revealed))
		for i := 0; i < len(revealed); i++ { // replace the _ with the letters
			if i == letter {
				underscoreWord[i] = revealed[i]
				if len(TabLetter) == 0 {
					H.Letter += revealed[i]
					H.Letter += " "
					TabLetter = append(TabLetter, revealed[i])
				} else {
					c := 0
					for _, k := range TabLetter {
						if k == revealed[i] {
							c++
						}
					}
					if 0 == c {
						H.Letter += revealed[i]
						H.Letter += " "
						TabLetter = append(TabLetter, revealed[i])
					}
				}
				for m := 0; m < len(revealed); m++ { // check if there is another similar letter
					if revealed[i] == revealed[m] {
						underscoreWord[m] = revealed[m]
					}
				}
			}
		}
	}
	for m := 0; m < len(Word); m++ { // transform a []string into a string
		Wordend += underscoreWord[m]
		Wordend += " "
	}
	return Wordend
}
