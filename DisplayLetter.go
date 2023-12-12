package Hangman

import "strings"

// checks whether the user has already used this character
func DisplayLetter(H *HangManData) bool {
	letter := strings.Split(H.Letter, " ") // character usde
	c := 0
	for i := 0; i < len(letter); i++ {
		if letter[i] != H.Input { // checks whether the character has been used
			c++
		}
	}
	if c != len(letter) { // or if is the first time
		return true
	}
	return false
}
