package Hangman

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// function to save everything in a text file
func Save(H *HangManData) {
	save, err := json.Marshal(H)
	if err != nil {
		os.Exit(1)
	}
	file, err := os.Create("save.txt")
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()
	_, err = file.Write(save)
	if err != nil {
		os.Exit(3)
	}
}

// function that retrieves and runs the program with the elements saved in the text file
func LoadGame(file string, hangman *HangManData) {
	load, err := ioutil.ReadFile(file)
	if err != nil {
		os.Exit(4)
	}
	err = json.Unmarshal(load, hangman)
	if err != nil {
		os.Exit(5)
	}
}
