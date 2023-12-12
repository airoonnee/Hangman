package Hangman

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// window to be displayed when the player wins
func Win(H *HangManData) {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextColor(tcell.ColorGreen).
		SetText("\n\n\n" + ":::   :::  ::::::::  :::    :::  :::       ::: ::::::::::: ::::    :::" + "\n" + ":+:   :+: :+:    :+: :+:    :+:  :+:       :+:     :+:     :+:+:   :+:" + "\n" + " +:+ +:+  +:+    +:+ +:+    +:+  +:+       +:+     +:+     :+:+:+  +:+" + "\n" + "  +#++:   +#+    +:+ +#+    +:+  +#+  +:+  +#+     +#+     +#+ +:+ +#+" + "\n" + "   +#+    +#+    +#+ +#+    +#+  +#+ +#+#+ +#+     +#+     +#+  +#+#+#" + "\n" + "   #+#    #+#    #+# #+#    #+#   #+#+# #+#+#      #+#     #+#   #+#+#" + "\n" + "   ###     ########   ########     ###   ###   ########### ###    ####" + "\n\n\n\n\n\n\n\n\n\n\n\n\n\n" + "The Word is :" + "\n\n" + AsciiArt(H.ToFind, H)).
		SetTextAlign(tview.AlignCenter).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// window to be displayed when the player loose
func Loose(H *HangManData) {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextColor(tcell.ColorRed).
		SetText("\n\n\n" + ":::   :::  ::::::::  :::    :::  :::         ::::::::   ::::::::   ::::::::  :::::::::: " + "\n" + ":+:   :+: :+:    :+: :+:    :+:  :+:        :+:    :+: :+:    :+: :+:    :+: :+:        " + "\n" + " +:+ +:+  +:+    +:+ +:+    +:+  +:+        +:+    +:+ +:+    +:+ +:+        +:+        " + "\n" + "  +#++:   +#+    +:+ +#+    +:+  +#+        +#+    +:+ +#+    +:+ +#++:++#++ +#++:++#   " + "\n" + "   +#+    +#+    +#+ +#+    +#+  +#+        +#+    +#+ +#+    +#+        +#+ +#+        " + "\n" + "   #+#    #+#    #+# #+#    #+#  #+#        #+#    #+# #+#    #+# #+#    #+# #+#        " + "\n" + "   ###     ########   ########   ##########  ########   ########   ########  ########## " + "\n\n\n\n\n" + H.Drow + "The Word is :" + "\n\n" + AsciiArt(H.ToFind, H)).
		SetTextAlign(tview.AlignCenter).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
