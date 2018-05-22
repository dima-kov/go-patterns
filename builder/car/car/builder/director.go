package builder

import "github.com/dima-kov/go-patterns/builder/car/car"

type CarBuilderDirector struct{}

func NewCarBuilderDirector() CarBuilderDirector {
	return CarBuilderDirector{}
}

func (cd *CarBuilderDirector) BuildSportCar(builder CarBuilder) {
	builder.reset()
	builder.setWheelsNumber(10)
	builder.setSeatsNumber(12)
	builder.setGlassType(car.RedGalss)
	builder.setGPS(true)
	builder.setEngineType(car.GasEngine)
}
