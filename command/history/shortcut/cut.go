package shortcut

import "github.com/dima-kov/go-patterns/command/history/command"

type CutShortcut struct {
	command command.Command
}

func (es *CutShortcut) Handle() {
	es.command.Execute()
}