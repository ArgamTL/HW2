package main

import "fmt"

type iBuilder interface {
	setBodyType()
	setEngineType()
	setGearBox()
	getCar() car
}

func getBuilder(builderType string) iBuilder {
	if builderType == "simple" {
		return &simpleBuilder{}
	}
	if builderType == "offroad" {
		return &offRoadBuilder{}
	}
	return nil
}

// normal

type simpleBuilder struct {
	bodyType   string
	engineType string
	gearbox    int
}

func newSimpleBuilder() *simpleBuilder {
	return &simpleBuilder{}
}

func (b *simpleBuilder) setBodyType() {
	b.bodyType = "sedan"
}

func (b *simpleBuilder) setEngineType() {
	b.engineType = "v"
}

func (b *simpleBuilder) setGearBox() {
	b.gearbox = 6
}

func (b *simpleBuilder) getCar() car {
	return car{
		engineType: b.engineType,
		bodyType:   b.bodyType,
		gearbox:    b.gearbox,
	}
}

type offRoadBuilder struct {
	bodyType   string
	engineType string
	gearbox    int
}

func newoffRoadBuilder() *offRoadBuilder {
	return &offRoadBuilder{}
}

func (b *offRoadBuilder) setBodyType() {
	b.bodyType = "Jeep"
}

func (b *offRoadBuilder) setEngineType() {
	b.engineType = "V"
}

func (b *offRoadBuilder) setGearBox() {
	b.gearbox = 5
}

func (b *offRoadBuilder) getCar() car {
	return car{
		engineType: b.engineType,
		bodyType:   b.bodyType,
		gearbox:    b.gearbox,
	}
}

//

type car struct {
	bodyType   string
	engineType string
	gearbox    int
}

//

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() car {
	d.builder.setEngineType()
	d.builder.setBodyType()
	d.builder.setGearBox()
	return d.builder.getCar()
}

func main() {
	simpleBuilder := getBuilder("simple")
	offRoadBuilder := getBuilder("offroad")

	director := newDirector(simpleBuilder)
	simpleCar := director.buildHouse()

	fmt.Printf("Simple car Engine type: %s\n", simpleCar.engineType)
	fmt.Printf("Simple car Body Type: %s\n", simpleCar.bodyType)
	fmt.Printf("Simple car Gears: %d\n", simpleCar.gearbox)

	director.setBuilder(offRoadBuilder)
	offRoadcar := director.buildHouse()

	fmt.Printf("\nOffroad car Engine type: %s\n", offRoadcar.engineType)
	fmt.Printf("Offroad car Body Type: %s\n", offRoadcar.bodyType)
	fmt.Printf("Offroad car  Gears: %d\n", offRoadcar.gearbox)
}
