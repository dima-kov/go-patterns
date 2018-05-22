package button_machine

import "fmt"

type buttonCoffeeMachine struct{}

func NewButtonCoffeeMachine() buttonCoffeeMachine {
	return buttonCoffeeMachine{}
}

func (oc buttonCoffeeMachine) PressButtonToFirst() {
	fmt.Println("Selected first in OLD MACHINE")
}

func (oc buttonCoffeeMachine) PressButtonToSecond() {
	fmt.Println("Selected second in OLD MACHINE")
}
