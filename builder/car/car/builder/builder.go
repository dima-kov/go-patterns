package builder

import "github.com/dima-kov/go-patterns/builder/car/car"

type carBuilder struct {
	car car.Car
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (cb *carBuilder) reset() {
	cb.car = car.NewCar()
}

func (cb *carBuilder) setSeatsNumber(number uint) {
	cb.car.SeatsNumber = number
}

func (cb *carBuilder) setWheelsNumber(number uint) {
	cb.car.WheelsNumber = number
}

func (cb *carBuilder) setGlassType(gType car.GlassType) {
	cb.car.GlassType = gType
}

func (cb carBuilder) setEngineType(engineType car.EngineType) {
	cb.car.EngineType = engineType
}

func (cb *carBuilder) setGPS(gps bool) {
	cb.car.GPS = gps
}

func (cb *carBuilder) GetResult() car.Car {
	return cb.car
}
