package components

import "fmt"

type Input struct {
	container   Container
	HelpMessage string
}

func (i *Input) ShowHelp() {
	if i.HelpMessage != "" {
		//	animated input to show message
		fmt.Println("[Input message] ", i.HelpMessage)
	} else {
		i.container.ShowHelp()
	}
}

func (i *Input) SetContainer(cont Container) {
	i.container = cont
}
