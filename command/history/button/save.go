package button

import (
	"github.com/dima-kov/go-patterns/command/history/command"
	"fmt"
)

type SaveButton struct {
	Label string

	Command command.Command
}

func (sb *SaveButton) Click() {
	fmt.Println("SaveButton clicked")
	sb.Command.Execute()
}