package main

import (
	"github.com/dima-kov/go-patterns/command/history/button"
	"github.com/dima-kov/go-patterns/command/history/command"
	"github.com/dima-kov/go-patterns/command/history/shortcut"
)

func main()  {
	saveCommand := command.SaveCommand{}
	saveButton := button.SaveButton{"Save", saveCommand}
	saveShortcut := shortcut.SaveShortcut{"ctrl-s", saveCommand}
	saveShortcut.Handle()
	saveButton.Click()
}
