package main

import (
	"fmt"
)

// CarModel — первая подсистема автомобиля, описывающая модель автомобиля.
type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Println(" CarModel - SetModel.")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Println(" CarEngine - SetEngine.")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Println(" CarBody - SetBody.")
}

type CarAccessories struct {
}

func NewCarAccessories() *CarAccessories {
	return &CarAccessories{}
}

func (c *CarAccessories) SetAccessories() {
	fmt.Println(" CarAccessories - SetAccessories")
}

// CarFacade описывает фасад автомобиля, который предоставляет упрощенный интерфейс для создания автомобиля.
type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

// NewCarFacade создает новый CarFacade.
func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

// MakeCar создает новый полный автомобиль.
func (c *CarFacade) MakeCar() {
	fmt.Println("==Creating a Car==")
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
	c.accessories.SetAccessories()
	fmt.Println("==Car creation is completed.==")
}

func main() {
	facade := NewCarFacade()
	facade.MakeCar()

}
