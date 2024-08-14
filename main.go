package main

import (
	"fmt"
	"math/rand"
	//"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/driver/desktop"
	"strings"

	"fyne.io/fyne/v2/widget"
)

type entry1 struct {
	widget.Entry
}

func newEntry1() *entry1 {
	e := &entry1{}
	e.ExtendBaseWidget(e)
	e.Entry.MultiLine = true
	e.Entry.Wrapping = fyne.TextWrapWord
	return e
}

func (e *entry1) TypedRune(r rune) {
	//fmt.Printf("entry1 tr: %v\n", r)
	e.Entry.TypedRune(r)
	//printCursorPosition(e)
}

func (e *entry1) TypedKey(ke *fyne.KeyEvent) {
	//fmt.Printf("entry1 ke: %v\n", ke)
	e.Entry.TypedKey(ke)

	if ke.Name == fyne.KeySpace {
		printCursorPosition(e)
	}
}

func printCursorPosition(e *entry1) {
	cursor := e.Entry.CursorColumn
	row := e.Entry.CursorRow
	text := e.Entry.Text
	lineText := strings.Split(text, "\n")[row]
	part1, part2 := text[:cursor], text[cursor:]
	fmt.Println("part 1: " + part1)
	fmt.Println("part 2: " + part2)
	fmt.Println(cursor, row, lineText)
	modifyWord := strings.Split(part1, " ")
	modifyWord = modifyWord[len(modifyWord)-1:]

	part1 = strings.TrimSuffix(part1, modifyWord[0])
	typoWord := typoWord(modifyWord[0])
	part1 += typoWord

	fmt.Println(part1 + " " + part2)

	e.Entry.SetText(part1 + " " + part2)
}

func typoWord(word string) string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	newWord := ""
	for i := 0; i < len(word); i++ {
		if rand.Float32() < 0.3 {
			newWord += string(charset[rand.Intn(len(charset))])
		} else {
			newWord += string(word[i])
		}
	}

	return newWord
}

func main() {
	a := app.New()
	w := a.NewWindow("Abc")

	editor1 := newEntry1()
	editor1.SetPlaceHolder("Start typing #1 ok...")

	content := container.NewStack(editor1)
	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 600))

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		fmt.Printf("window canvas: %v\n", ke.Name)
	})

	w.ShowAndRun()
}
