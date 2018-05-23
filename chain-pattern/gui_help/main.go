package main

import "github.com/dima-kov/go-patterns/chain-pattern/gui_help/components"

func main() {
	body := components.Body{HelpMessage:"xscds"}
	div := components.Div{HelpMessage:"sdcd"}
	nameInput := components.Input{HelpMessage:"Name, please"}
	okButton := components.Button{HelpMessage:"Submit to send OK"}

	div.AddItems(&nameInput, &okButton)
	body.AddItems(&div)

	okButton.ShowHelp()
}
