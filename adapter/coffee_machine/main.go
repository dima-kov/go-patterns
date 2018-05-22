package main

type HumanMakeCoffeeInterface interface {
	touchFirst()
	touchSecond()
}

func main()  {
	touchMachine := NewTouchMachine()
	oldMachine := NewButtonCoffeeMachine()

	touchMachine.touchFirst()
	touchMachine.touchSecond()

	touchAdapterForOld := NewTouchCoffeeMachineAdapter(oldMachine)
	touchAdapterForOld.touchFirst()
	touchAdapterForOld.touchSecond()
}
