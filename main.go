package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("tbd")

	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Start typing...")


	content := container.NewStack(editor)
	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}