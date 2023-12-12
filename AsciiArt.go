package Hangman

import (
	"bufio"
	"fmt"
	"os"
)

// checks whether the user has entered any arguments after words.txt
func Ascii() bool {
	if len(os.Args) == 2 {
		return false
	}
	return true
}

func AsciiArt(Word string, H *HangManData) string {

	if Ascii() {
		if H.Ascii == "Ascii/shadow.txt" || H.Ascii == "Ascii/standard.txt" || H.Ascii == "Ascii/thinkertoy.txt" {
			Ascii := ""
			for i := 0; i < 9; i++ {
				for _, RuneLetter := range Word {
					file, err := os.Open(H.Ascii)
					if err != nil {
						fmt.Println("Error opening file:", err)
					}
					// Create a scanner to read the contents of the file line by line
					scanner := bufio.NewScanner(file)
					Line := 0
					for scanner.Scan() {
						if Line == ((int(RuneLetter)-97)*9)+586+i {
							Ascii += scanner.Text()
						}
						Line++
					}
				}
				Ascii += "\n"
			}
			return Ascii
		} else if os.Args[2] == "--letterFile" && (os.Args[3] == "Ascii/shadow.txt" || os.Args[3] == "Ascii/standard.txt" || os.Args[3] == "Ascii/thinkertoy.txt") {
			Ascii := ""
			for i := 0; i < 9; i++ {
				for _, RuneLetter := range Word {
					file, err := os.Open(AsciiFile(H))
					if err != nil {
						fmt.Println("Error opening file:", err)
					}
					// Create a scanner to read the contents of the file line by line
					scanner := bufio.NewScanner(file)
					Line := 0
					for scanner.Scan() {
						if Line == ((int(RuneLetter)-97)*9)+586+i {
							Ascii += scanner.Text()
						}
						Line++
					}
				}
				Ascii += "\n"
			}
			return Ascii
		}
	}
	return Word
}
