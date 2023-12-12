package Hangman

import "strings"

// checks whether the character entered is a word or a letter
func Check(H *HangManData) string {
	Input := strings.Split(H.Input, "")          // character enter
	WordFind := strings.Split(H.ToFind, "")      // Word to find
	underscoreWord := strings.Split(H.Word, " ") // hide wordr
	var Wordend string

	if len(Input) == 1 {
		if Input[0] >= "a" && Input[0] <= "z" { // check if it's a lowercase letter
			DisplayLetter(H)

			count := 0
			for i := 0; i < len(WordFind); i++ { //

				if WordFind[i] == Input[0] { // checks whether the letter is in the word
					underscoreWord[i] = Input[0]

					count += 1
				}
			}
			if count == 0 { // the letter is not in the word
				H.Attempts += 1
				False(H)

				return H.Word
			}
			for m := 0; m < len(WordFind); m++ { // transform a []string into a string
				Wordend += underscoreWord[m]
				Wordend += " "
			}
			H.Word = Wordend
		}
	} else {
		c := 0
		for a := 0; a < len(Input); a++ {
			if Input[a] >= "a" && Input[a] <= "z" {
				if WordFind[a] == Input[a] {
					c++
				} else {
					H.Attempts += 2
					False(H)
					return H.Word
				}
				if c == len(WordFind) {
					H.Word = H.ToFind
				}
			}
		}
	}
	return H.Word
}

// check if the word to guess contains _.
func VerificationForWin(H *HangManData) bool {
	for _, i := range H.Word {
		if i == '_' {
			return false
		}
	}
	return true
}
