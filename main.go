package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window
var filePath string
var saved bool
var debug bool

type editor struct {
	widget.Entry
}

func newEditor() *editor {
	e := &editor{}
	e.ExtendBaseWidget(e)
	e.Entry.MultiLine = true
	e.Entry.Wrapping = fyne.TextWrapOff
	return e
}

func (e *editor) TypedKey(ke *fyne.KeyEvent) {
	e.Entry.TypedKey(ke)

	//fmt.Println(ke.Physical.ScanCode)

	if _, ok := nonTextChangeKeys[ke.Physical.ScanCode]; ok {
		return
	}

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

func sorryDialogContent(content string) fyne.CanvasObject {
	var label *widget.RichText

	if content == "filenotfound" {
		label = widget.NewRichTextFromMarkdown("The requested file \"" + filepath.Base(filePath) + "\" wasn't found. Make sure you didn't make a typo.")
	} else {
		label = widget.NewRichTextFromMarkdown("Sorry, you can't " + content + " files on web or mobile. Try downloading the desktop version on [Github](https://github.com/radeeyate/TypoTolerantTextEditor)!")
	}

	return container.NewVBox(label)
}

func saveFile(e *editor) {
	if app.New().Driver().Device().IsBrowser() || app.New().Driver().Device().IsMobile() {
		dialog.NewCustom("Download for Desktop", "OK", sorryDialogContent("save"), w).Show()
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
		dialog.NewCustom("Download for Desktop", "OK", sorryDialogContent("open"), w).Show()
	} else {
		fmt.Println(saved)
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

		content, err := os.ReadFile(file.URI().Path())
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}

		w.Canvas().Focus(e)
		e.Entry.SetText(string(content))
		saved = true
		filePath = file.URI().Path()
		w.SetTitle("Typo Tolerant Text Editor - " + file.URI().Name())
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
		result := ""
		replacementWord := replacements[rand.Intn(len(replacements))]
		for i := 0; i < len(word); i++ {
			if unicode.IsUpper(rune(word[i])) {
				result += strings.ToUpper(string(replacementWord[i]))
			} else {
				result += string(replacementWord[i])
			}
		}

		return result
	}

	var typoProbability float32
	var changeLimit int
	switch {
	case len(word) < 5:
		typoProbability = 0.3
		changeLimit = 3
	case len(word) >= 5:
		typoProbability = 0.2
		changeLimit = 4
	case len(word) >= 7:
		typoProbability = 0.15
		changeLimit = 5
	case len(word) >= 10:
		typoProbability = 0.05
		changeLimit = 6
	}

	result := ""

	if rand.Float32() < 0.1 && len(word) >= 4 { // 10% chance to flip characters
		wordRune := []rune(word)
		wordRune[0], wordRune[1] = wordRune[1], wordRune[0]
		result = string(wordRune)

		fmt.Println(result)
		return result
	}

	for i := 0; i < len(word); i++ {
		if rand.Float32() < typoProbability {
			changeCharacter := string(word[i])

			if changeLimit == 0 {
				result += string(word[i])
			}

			if replacements, ok := keyboardMap[strings.ToLower(changeCharacter)]; ok {
				characterToAdd := replacements[rand.Intn(len(replacements))]

				fmt.Println(changeCharacter)

				if unicode.IsUpper(rune(changeCharacter[0])) {
					characterToAdd = strings.ToUpper(characterToAdd)
				}

				result += characterToAdd
				changeLimit--
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
	ab := a.NewWindow("About Typo Tolerant Text Editor")
	ab.Resize(fyne.NewSize(500, 300))

	saved = true

	icon := canvas.NewImageFromResource(a.Metadata().Icon)
	icon.Resize(fyne.NewSize(50, 50))
	icon.FillMode = canvas.ImageFillOriginal

	version := a.Metadata().Version
	sourceLink, _ := url.Parse("https://github.com/radeeyate/TypoTolerantTextEditor")

	aboutInfo := container.NewVBox(
		icon,
		container.NewVBox(
			widget.NewLabelWithStyle("Typo Tolerant Text Editor\n"+version, fyne.TextAlignCenter, fyne.TextStyle{}),
			widget.NewLabelWithStyle("The text editor with built-in typos.", fyne.TextAlignCenter, fyne.TextStyle{}),
			widget.NewHyperlinkWithStyle("Source Code", sourceLink, fyne.TextAlignCenter, fyne.TextStyle{}),
		),
	)
	ab.SetContent(aboutInfo)

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

	debug = *flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if flag.NArg() > 0 {
		filePath = flag.Arg(0)
		fmt.Println(filePath)

		if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
			fmt.Println("exists!")
			content, err := os.ReadFile(filePath)

			if err != nil {
				dialog.NewError(err, w).Show()
			} else {
				editor.Entry.SetText(string(content))
				w.SetTitle("Typo Tolerant Text Editor - " + filepath.Base(filePath))
			}
		} else {
			dialog.NewCustom("Error", "OK", sorryDialogContent("filenotfound"), w).Show()
		}
	}

	if !a.Driver().Device().IsBrowser() && !a.Driver().Device().IsMobile() {
		w.SetMainMenu(makeMenu(editor, ab))
	}

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func makeMenu(e *editor, ab fyne.Window) *fyne.MainMenu {
	saveItem := fyne.NewMenuItem("Save (Ctrl+S)", func() {
		saveFile(e)
	})
	openItem := fyne.NewMenuItem("Open (Ctrl+O)", func() {
		openFileSaveCheck(e)
	})
	fileMenu := fyne.NewMenu("File", saveItem, openItem)

	aboutItem := fyne.NewMenuItem("About", func() {
		ab.Show()
	})
	helpMenu := fyne.NewMenu("Help", aboutItem)

	return fyne.NewMainMenu(fileMenu, helpMenu)
}
