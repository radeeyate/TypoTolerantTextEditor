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

	//	"unicode"

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
var currentProbability float32

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

func (e *editor) KeyUp(ke *fyne.KeyEvent) {
	e.Entry.KeyUp(ke)

	if _, ok := nonTextChangeKeys[ke.Physical.ScanCode]; ok {
		return
	} else {
		if !strings.HasPrefix(w.Title(), "*") {
			w.SetTitle("*" + w.Title())
			saved = false
		}

		if filePath == "" && e.Entry.Text == "" {
			w.SetTitle(strings.TrimPrefix(w.Title(), "*"))
		}

		if ke.Name == fyne.KeySpace {
			currentProbability = 0.15
		}

		if _, ok := keyboardMap[strings.ToLower(string(ke.Name))]; !ok { // character is not in replacement map
			return
		}

		modifyText(e)
	}
}

func modifyText(e *editor) {
	cursorColumn := e.Entry.CursorColumn
	cursorRow := e.Entry.CursorRow
	entryText := e.Entry.Text
	textLines := strings.Split(entryText, "\n")

	if e.Entry.Text == "" {
		return
	}

	if cursorRow >= len(textLines) {
		fmt.Println("Error: Row out of bounds")
		return
	}

	currentLine := textLines[cursorRow]
	textBeforeCursor, textAfterCursor := currentLine[:cursorColumn], currentLine[cursorColumn:]
	typedCharacter := currentLine[len(textBeforeCursor)-1:]

	if rand.Float32() < currentProbability {
		if replacements, ok := keyboardMap[strings.ToLower(typedCharacter)]; ok {
			characterToAdd := replacements[rand.Intn(len(replacements))]

			if unicode.IsUpper(rune(typedCharacter[0])) {
				characterToAdd = strings.ToUpper(characterToAdd)
			}
			
			e.Entry.SetText(textBeforeCursor[:len(textBeforeCursor)-1] + characterToAdd + textAfterCursor)
			if currentProbability < 0.35 {
				currentProbability += (0.05 + rand.Float32()*(0.1-0.05))

				if debug {
					fmt.Println(currentProbability)
				}
			}
		}
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

func main() {
	a := app.New()
	w = a.NewWindow("Typo Tolerant Text Editor")
	ab := a.NewWindow("About Typo Tolerant Text Editor")
	ab.Resize(fyne.NewSize(500, 300))

	saved = true
	currentProbability = 0.15

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

	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	if flag.NArg() > 0 {
		filePath = flag.Arg(0)

		if debug {
			fmt.Println(filePath)
		}

		if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
			if debug {
				fmt.Println("exists!")
			}
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
