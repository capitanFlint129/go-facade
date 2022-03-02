package builder

import "architectural-patterns-in-go/pkg/product"

type carBuilderInterface interface {
	GetResult() product.Product

	setSeats(seatsNumber int)
	setEngine(enginePower int)
	setTripComputer(tripComputerModel string)
	setGps(gpsModel string)
}

type carBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	carCreator        CarCreator // functor - скрываем конструктор за параметром и не привязываемся к конкретному
}

func (c *carBuilder) GetResult() product.Product {
	return c.carCreator(
		c.seatsNumber,
		c.enginePower,
		c.tripComputerModel,
		c.gpsModel,
	)
}

func (c *carBuilder) setSeats(seatsNumber int) {
	c.seatsNumber = seatsNumber
}

func (c *carBuilder) setEngine(enginePower int) {
	c.enginePower = enginePower
}

func (c *carBuilder) setTripComputer(tripComputerModel string) {
	c.tripComputerModel = tripComputerModel
}

func (c *carBuilder) setGps(gpsModel string) {
	c.gpsModel = gpsModel
}

func NewCarBuilder(carCreator CarCreator) Builder {
	return &carBuilder{
		carCreator: carCreator,
	}
}