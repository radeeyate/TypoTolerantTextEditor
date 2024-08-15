package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window

type editor struct {
	widget.Entry
}

var windowFile os.File
var filePath string

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

	if filePath == "" && e.Entry.Text == "" {
		w.SetTitle(strings.TrimPrefix(w.Title(), "*"))
	}

	if ke.Name == fyne.KeySpace {
		modifyText(e)
	}
}

func (e *editor) TypedShortcut(s fyne.Shortcut) {
	if shortcut, ok := s.(*desktop.CustomShortcut); ok {
		if shortcut.KeyName == fyne.KeyS && shortcut.Modifier == fyne.KeyModifierControl {
			saveFile(e)
			return
		} else if shortcut.KeyName == fyne.KeyO && shortcut.Modifier == fyne.KeyModifierControl {
			openFileSaveCheck(e)
			return
		}
	}
}

func saveFile(e *editor) {
	if app.New().Driver().Device().IsBrowser() || app.New().Driver().Device().IsMobile() {
		dialog.NewInformation("Error", "Sorry, you can't save files on web or mobile. Try downloading the desktop version on GitHub!", w).Show()
	} else {
		if filePath == "" {
			dialog.NewFileSave(func(file fyne.URIWriteCloser, err error) {

				if err != nil {
					dialog.NewError(err, w).Show()
					return
				}

				if file == nil {
					return
				}

				defer file.Close()
				file.Write([]byte(e.Text))
				filePath = file.URI().Path()
				w.SetTitle("Typo Tolerant Text Editor - " + file.URI().Name())
				saved = true
			}, w).Show()

		} else {
			err := os.WriteFile(filePath, []byte(e.Text), 0644)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}

			w.SetTitle(strings.TrimPrefix(w.Title(), "*"))
		}
	}
}

func openFileSaveCheck(e *editor) {
	if app.New().Driver().Device().IsBrowser() || app.New().Driver().Device().IsMobile() {
		dialog.NewInformation("Error", "Sorry, you can't open files on web or mobile. Try downloading the desktop version on GitHub!", w).Show()
	} else {
		if !saved {
			dialog.NewConfirm("Save", "You have unsaved changes. Are you sure you want to open a new file?", func(open bool) {
				if open {
					openFile(e)
				}
			}, w).Show()
		} else {
			openFile(e)
		}
	}
}

func openFile(e *editor) {
	dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}

		if file == nil {
			return
		}

		if _, err := os.Stat(file.URI().Path()); err != nil {
			dialog.NewError(err, w).Show()
			return
		}
		
		content, err := ioutil.ReadFile(file.URI().Path())
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}

		w.Canvas().Focus(e)
		e.Entry.SetText(string(content))
		saved = true
		filePath = file.URI().Path()
		w.SetTitle(strings.TrimPrefix(w.Title(), "*") + " - " + file.URI().Name())
	}, w).Show()
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
		typoProbability = 0.2
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

	w.SetContent(container.NewStack(editor))

	w.SetCloseIntercept(func() {
		if !saved {
			dialog.NewConfirm("Exit", "Are you sure you want to exit? Any unsaved changes will be lost.", func(close bool) {
				if close {
					os.Exit(0)
				}
			}, w).Show()
		} else {
			os.Exit(0)
		}
	})
	
	w.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyS,
		Modifier: fyne.KeyModifierControl,
	}, func(shortcut fyne.Shortcut) { saveFile(editor) })

	w.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyO,
		Modifier: fyne.KeyModifierControl,
	}, func(shortcut fyne.Shortcut) { openFileSaveCheck(editor) })

	if !a.Driver().Device().IsBrowser() && !a.Driver().Device().IsMobile() {
		w.SetMainMenu(makeMenu(a, w, editor))
	}

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window, e *editor) *fyne.MainMenu {
	saveItem := fyne.NewMenuItem("Save (Ctrl+S)", func() {
		saveFile(e)
	})
	openItem := fyne.NewMenuItem("Open (Ctrl+O)", func() {
		openFileSaveCheck(e)
	})
	fileMenu := fyne.NewMenu("File", saveItem, openItem)

	return fyne.NewMainMenu(fileMenu)
}
