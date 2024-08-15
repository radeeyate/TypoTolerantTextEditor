package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"os"
	"strings"
)

var w fyne.Window

type editor struct {
	widget.Entry
}

var windowFile os.File
var fileName string

var saved bool

func newEditor() *editor {
	e := &editor{}
	e.ExtendBaseWidget(e)
	e.Entry.MultiLine = true
	e.Entry.Wrapping = fyne.TextWrapOff
	return e
}

func (e *editor) TypedKey(ke *fyne.KeyEvent) {
	e.Entry.TypedKey(ke)

	if !strings.HasPrefix(w.Title(), "*") {
		w.SetTitle("*" + w.Title())
		saved = false
	}

	if ke.Name == fyne.KeySpace {
		modifyText(e)
	}
}

func (e *editor) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		e.Entry.TypedShortcut(s)
		return
	}

	if s.ShortcutName() == "CustomDesktop:Control+S" {
		if app.New().Driver().Device().IsBrowser() || app.New().Driver().Device().IsMobile() {
			dialog.NewInformation("Error", "Sorry, you can't save files on web or mobile. Try downloading the desktop version!", w).Show()
		} else {
			saver := dialog.NewFileSave(func(file fyne.URIWriteCloser, err error) {
				if err != nil {
					return
				}
				file.Write([]byte(e.Text))
				fileName = file.URI().Name()
				w.SetTitle("Typo Tolerant Text Editor - " + fileName)
				saved = true
			}, w)
			saver.Show()
		}
	}

}

func modifyText(e *editor) {
	cursorColumn := e.Entry.CursorColumn
	cursorRow := e.Entry.CursorRow
	entryText := e.Entry.Text

	lines := strings.Split(entryText, "\n")
	if cursorRow >= len(lines) {
		fmt.Println("Error: Row out of bounds")
		return
	}

	currentLine := lines[cursorRow]
	textBeforeCursor, textAfterCursor := currentLine[:cursorColumn], currentLine[cursorColumn:]

	// debugging prints
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		fmt.Println("Text before cursor: " + textBeforeCursor)
		fmt.Println("Text after cursor: " + textAfterCursor)
		fmt.Println(cursorColumn, cursorRow, currentLine)
	}

	words := strings.Split(textBeforeCursor, " ")
	if len(words) == 0 {
		return // no words to modify
	}
	lastWord := words[len(words)-1]

	textBeforeCursor = strings.TrimSuffix(textBeforeCursor, lastWord)
	correctedWord := introduceTypo(lastWord)
	textBeforeCursor += correctedWord

	if len(textAfterCursor) == 0 {
		lines[cursorRow] = textBeforeCursor
	} else {
		lines[cursorRow] = textBeforeCursor + textAfterCursor
	}

	updatedText := strings.Join(lines, "\n")
	e.Entry.SetText(updatedText)
}

func introduceTypo(word string) string {
	if replacements, ok := wordReplacements[word]; ok {
		return replacements[rand.Intn(len(replacements))]
	}

	var typoProbability float32
	switch {
	case len(word) < 5:
		typoProbability = 0.3
	case len(word) > 5:
		typoProbability = 0.15
	}

	result := ""
	for i := 0; i < len(word); i++ {
		if rand.Float32() < typoProbability {
			if replacements, ok := keyboardMap[string(word[i])]; ok {
				result += replacements[rand.Intn(len(replacements))]
			} else {
				result += string(word[i])
			}
		} else {
			result += string(word[i])
		}
	}

	return result
}

func main() {
	a := app.New()
	w = a.NewWindow("Typo Tolerant Text Editor")

	saved = true

	editor := newEditor()
	editor.SetPlaceHolder("Start typing here...")

	content := container.NewStack(editor)
	w.SetContent(content)

	w.SetMainMenu(makeMenu(a, w))

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = false
	newItem.ChildMenu = fyne.NewMenu("File", checkedItem)

	return fyne.NewMainMenu(newItem.ChildMenu)
}
