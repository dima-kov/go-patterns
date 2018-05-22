package car

type GlassType string
type EngineType string

const (
	WhiteGlass GlassType = "white"
	BlackGlass GlassType = "black"
	RedGalss   GlassType = "red"

	OilEngine    EngineType = "oil"
	PetrolEngine EngineType = "petrol"
	GasEngine    EngineType = "gas"
)

type Car struct {
	SeatsNumber  uint
	WheelsNumber uint
	GlassType    GlassType
	EngineType   EngineType
	GPS          bool
}

func NewCar() Car {
	return Car{}
}
