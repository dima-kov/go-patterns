package button_machine

import "github.com/dima-kov/go-patterns/adapter/coffee_machine/machine_interface"

type touchCoffeeMachineAdapter struct {
	buttonMachine buttonCoffeeMachine
}

func NewTouchCoffeeMachineAdapter(machine buttonCoffeeMachine) machine_interface.HumanMakeCoffeeInterface {
	return touchCoffeeMachineAdapter{machine}
}

func (ca touchCoffeeMachineAdapter) TouchFirst() {
	ca.buttonMachine.PressButtonToFirst()
}

func (ca touchCoffeeMachineAdapter) TouchSecond() {
	ca.buttonMachine.PressButtonToSecond()
}
