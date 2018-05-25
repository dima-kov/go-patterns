package command

import "fmt"

type SaveCommand struct {}

func (sv SaveCommand) Execute() {
	//make save
	fmt.Println("SAVING editor")
}
