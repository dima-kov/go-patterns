package touch_machine

import (
	"fmt"
	"github.com/dima-kov/go-patterns/adapter/coffee_machine/machine_interface"
)

type touchMachine struct{}

func NewTouchMachine() machine_interface.HumanMakeCoffeeInterface {
	return touchMachine{}
}

func (nm touchMachine) TouchFirst() {
	fmt.Println("Selected first in NEW MACHINE")
}

func (nm touchMachine) TouchSecond() {
	fmt.Println("Selected second in NEW MACHINE")
}
