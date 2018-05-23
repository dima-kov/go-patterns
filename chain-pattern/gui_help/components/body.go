package components

import "fmt"

type Body struct {
	HelpMessage string

	children []Component
}

func (b *Body) ShowHelp() {
	if b.HelpMessage != "" {
		fmt.Println("[Body message]", b.HelpMessage)
	}
}

func (b *Body) SetContainer(cont Container) {

}

func (b *Body) AddItems(items ...Component) {
	b.children = append(b.children, items...)
}
