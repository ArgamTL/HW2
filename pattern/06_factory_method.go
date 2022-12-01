package main

import "fmt"

type iCar interface {
	setName(name string)
	setSpeed(speed int)
	getName() string
	getSpeed() int
}

type car struct {
	name  string
	speed int
}

func (g *car) setName(name string) {
	g.name = name
}

func (g *car) getName() string {
	return g.name
}

func (g *car) setSpeed(speed int) {
	g.speed = speed
}

func (g *car) getSpeed() int {
	return g.speed
}

type sportcar struct {
	car
}

func newSportCar() iCar {
	return &sportcar{
		car: car{
			name:  "Spotr car",
			speed: 300,
		},
	}
}

type offroad struct {
	car
}

func newOffroadCar() iCar {
	return &offroad{
		car: car{
			name:  "Offroad car",
			speed: 90,
		},
	}
}

func getCar(carType string) (iCar, error) {
	if carType == "sportcar" {
		return newSportCar(), nil
	}
	if carType == "offroad" {
		return newOffroadCar(), nil
	}
	return nil, fmt.Errorf("Wrong car type passed")
}

func main() {
	sportcar, _ := getCar("sportcar")
	offroad, _ := getCar("offroad")
	printDetails(sportcar)
	printDetails(offroad)
}

func printDetails(g iCar) {
	fmt.Printf("Car: %s", g.getName())
	fmt.Println()
	fmt.Printf("Speed: %d km/h", g.getSpeed())
	fmt.Println()
}
