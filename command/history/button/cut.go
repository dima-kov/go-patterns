package button

import (
	"github.com/dima-kov/go-patterns/command/history/command"
	"fmt"
)

type CutButton struct {
	label string

	command command.Command
}

func (sb *CutButton ) Click() {
	fmt.Println("CutButton clicked")
	sb.command.Execute()
}