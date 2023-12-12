package Hangman

import (
	"fmt"
	"os"
)

func File() string {
	// Check that the filename argument has been supplied
	if len(os.Args) < 2 {
		fmt.Println("the file name is not provided")
	}
	// Get the file name from the command line arguments
	fileName := os.Args[1]
	return fileName
}

func AsciiFile(H *HangManData) string {
	// Check that the filename argument has been supplied
	if len(os.Args) < 3 {
		fmt.Println("the name of the ASCII Art file is not supplied")
	}
	// Get the file name from the command line arguments
	fileName := os.Args[3]
	H.Ascii = fileName
	return fileName
}
