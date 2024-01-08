package Hangman

// function used when the given letter is not in the word to be guessed
// func False(H *HangManData) string {
// 	Drow := ""
// 	file, err := os.Open("hangman.txt")
// 	if err != nil {
// 		log.Fatalf("Error when opening file: %s", err)
// 	}
// 	fileScanner := bufio.NewScanner(file)
// 	Line := 0
// 	for fileScanner.Scan() {
// 		if H.Attempts <= 0 {
// 			if Line >= 72 && Line <= 80 {
// 				Drow += fileScanner.Text() + "\n"
// 			}
// 		} else if Line >= H.HangmanPositions[H.Attempts] && Line < H.HangmanPositions[H.Attempts-1] && H.Attempts <= 9 {
// 			Drow += fileScanner.Text() + "\n"
// 		}
// 		Line++
// 	}
// 	H.Drow = Drow
// 	return H.Drow
// }
