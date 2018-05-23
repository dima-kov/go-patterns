package components

import "fmt"

type Button struct {
	container Container

	HelpMessage string
}

func (b *Button) ShowHelp() {
	if b.HelpMessage != "" {
		// Some fantastic method to show tooltip message
		fmt.Println("[Message] ", b.HelpMessage)
	} else {
		b.container.ShowHelp()
	}
}

func (b *Button) SetContainer(cont Container) {
	b.container = cont
}