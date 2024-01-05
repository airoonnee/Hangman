package Hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// choose a word at random
func Word() string {
	// Open the file for reading
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	// Create a scanner to read the contents of the file line by line
	scanner := bufio.NewScanner(file)
	// Store words in a slice
	var words []string
	for scanner.Scan() {
		line := scanner.Text()
		wordsInLine := strings.Fields(line)
		words = append(words, wordsInLine...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
	rand.Seed(time.Now().UnixNano())
	// Generate a random index to select a word
	randomIndex := rand.Intn(len(words))
	// Select the word
	randomWord := words[randomIndex]
	return randomWord
}
