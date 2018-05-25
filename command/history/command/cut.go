package command

import "fmt"

type CutCommand struct {}

func (sv *CutCommand ) Execute() {
	//make save
	fmt.Println("Cut selected text from editor")
}
