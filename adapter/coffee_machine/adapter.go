package main

type touchCoffeeMachineAdapter struct {
	oldMachine buttonCoffeeMachine
}

func NewTouchCoffeeMachineAdapter(machine buttonCoffeeMachine) touchCoffeeMachineAdapter {
	return touchCoffeeMachineAdapter{machine}
}

func (ca *touchCoffeeMachineAdapter) touchFirst() {
	ca.oldMachine.pressButtonToFirst()
}

func (ca *touchCoffeeMachineAdapter) touchSecond() {
	ca.oldMachine.pressButtonToSecond()
}