package Hangman

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// display of the game and each small window
// and the launch of each function when the player enters a character
func Forme(H *HangManData) {
	Question := "Enter letter tiny : "
	Hangman := tview.NewTextView().
		SetText(H.Drow).
		SetTextAlign(tview.AlignCenter)

	Hangman.SetBorder(true).
		SetTitle("Hangman").
		SetBorder(true)

	Mot := tview.NewTextView().
		SetText(AsciiArt(H.Word, H)).
		SetTextAlign(tview.AlignCenter)

	Mot.SetBorder(true).
		SetTitle("Mot").
		SetBorder(true)

	lettres := tview.NewTextView().
		SetText(H.Letter).
		SetTextAlign(tview.AlignCenter)

	lettres.SetBorder(true).
		SetTitle("Letters").
		SetBorder(true)

	app := tview.NewApplication()
	app.EnableMouse(true)

	WriteLetter := tview.NewInputField().
		SetLabel(Question).
		SetFieldWidth(10)

	WriteLetter.SetDoneFunc(func(key tcell.Key) {
	}).
		SetBorder(true).
		SetTitle("write").
		SetBorder(true)

	button := tview.NewForm().
		AddButton("Quit", func() {
			app.Stop()
		})

		// when the player enters a character
	WriteLetter.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			H.Input = WriteLetter.GetText()
			Input := strings.Split(H.Input, "")
			Word := strings.Split(H.ToFind, "")
			if H.Input == "STOP" {
				Save(H)
				app.Stop()
			}
			if len(Input) == 1 || len(Input) == len(Word) { // check whether a letter or word is entered
				c := 0
				for i := 0; i < len(Input); i++ {
					if Input[i] >= "a" && Input[i] <= "z" {
						c++
						if c == len(Input) {
							DisplayLetter(H)
							if !DisplayLetter(H) {
								Check(H)
								Mot.SetText(AsciiArt(H.Word, H))
								WriteLetter.SetText("")
								WriteLetter.SetLabel(Question)
								Hangman.SetText(H.Drow)
								H.Letter += H.Input
								H.Letter += " "
								lettres.SetText(H.Letter)
								VerificationForWin(H)
								if VerificationForWin(H) {
									Win(H)
								}
								if H.Attempts >= 10 {
									Loose(H)
								}
							} else {
								WriteLetter.SetLabel("This character has been used before, retry :")
								WriteLetter.SetText("")
							}
						}
					} else {
						WriteLetter.SetLabel("This is not a lower case letter, retry :")
						WriteLetter.SetText("")
					}
				}
			} else { // if it enters a word that is not the same size as the word to be searched for
				c := 0
				for i := 0; i < len(Input); i++ {
					if Input[i] >= "a" && Input[i] <= "z" {
						c++
						if c == len(Input) {
							DisplayLetter(H)
							if !DisplayLetter(H) {
								H.Attempts += 2
								False(H)
								Mot.SetText(AsciiArt(H.Word, H))
								WriteLetter.SetText("")
								Hangman.SetText(H.Drow)
								H.Letter += H.Input
								H.Letter += " "
								lettres.SetText(H.Letter)
								if H.Attempts >= 10 {
									Loose(H)
								}
							} else {
								WriteLetter.SetLabel("This character has been used before, retry :")
								WriteLetter.SetText("")
							}
						}
					} else {
						WriteLetter.SetLabel("This is not a lower case letter, retry :")
						WriteLetter.SetText("")

					}
				}
			}
		}
	})

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(Hangman, 0, 6, false).
			AddItem(button, 0, 1, false), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(Mot, 0, 1, false).
			AddItem(lettres, 0, 1, false).
			AddItem(WriteLetter, 0, 1, false), 0, 2, false)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
