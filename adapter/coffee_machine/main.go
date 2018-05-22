package main

import (
	"github.com/dima-kov/go-patterns/adapter/coffee_machine/button_machine"
	"github.com/dima-kov/go-patterns/adapter/coffee_machine/touch_machine"
)

func main() {
	touchMachine := touch_machine.NewTouchMachine()
	oldMachine := button_machine.NewButtonCoffeeMachine()

	touchMachine.TouchFirst()
	touchMachine.TouchSecond()

	touchAdapterForButton := button_machine.NewTouchCoffeeMachineAdapter(oldMachine)
	touchAdapterForButton.TouchFirst()
	touchAdapterForButton.TouchSecond()
}
