package shortcut

import (
	"github.com/dima-kov/go-patterns/command/history/command"
	"fmt"
)

type SaveShortcut struct {
	Keys    string
	Command command.Command
}

func (ss *SaveShortcut) Handle() {
	fmt.Println("SaveShortcut binding")
	ss.Command.Execute()
}
