package builder

import "github.com/dima-kov/go-patterns/builder/car/car"

type CarBuilder interface {
	reset()
	setSeatsNumber(number uint)
	setWheelsNumber(number uint)
	setGlassType(gType car.GlassType)
	setEngineType(engineType car.EngineType)
	setGPS(gps bool)
	GetResult() car.Car
}
