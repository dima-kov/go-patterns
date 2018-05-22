package main

import (
	"fmt"
	"github.com/dima-kov/go-patterns/builder/car/car/builder"
)

func main() {
	builderDirector := builder.NewCarBuilderDirector()

	carBuilder := builder.NewCarBuilder()
	builderDirector.BuildSportCar(carBuilder)
	car := carBuilder.GetResult()
	fmt.Println("CAR:", car)
}
