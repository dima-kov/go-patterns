package main

import "fmt"

type buttonCoffeeMachine struct {}

func NewButtonCoffeeMachine() buttonCoffeeMachine {
	return buttonCoffeeMachine{}
}


func (oc buttonCoffeeMachine) pressButtonToFirst() {
	fmt.Println("Selected first in OLD MACHINE")
}

func (oc buttonCoffeeMachine) pressButtonToSecond() {
	fmt.Println("Selected second in OLD MACHINE")
}