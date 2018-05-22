package main

import "fmt"

type touchMachine struct {}

func NewTouchMachine() HumanMakeCoffeeInterface {
	return touchMachine{}
}

func (nm touchMachine) touchFirst() {
	fmt.Println("Selected first in NEW MACHINE")
}

func (nm touchMachine) touchSecond() {
	fmt.Println("Selected second in NEW MACHINE")
}