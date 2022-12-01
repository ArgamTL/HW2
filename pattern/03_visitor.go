package main

import "fmt"

type control interface {
	getType() string
	accept(visitor)
}

type gearshift struct {
	gear string
}

func (c *gearshift) accept(v visitor) {
	v.visitForGearShift(c)
}

func (c *gearshift) getType() string {
	return "Gear: "
}

// //
type steeringWheel struct {
	steering string
}

func (s *steeringWheel) accept(v visitor) {
	v.visitForSteeringWheel(s)
}

func (s *steeringWheel) getType() string {
	return "Steering Wheel: "
}

// vis

type visitor interface {
	visitForSteeringWheel(*steeringWheel)
	visitForGearShift(*gearshift)
}

type smartAuto struct {
	selfdriving string
}

func (t *smartAuto) visitForSteeringWheel(s *steeringWheel) {
	fmt.Println("Add selfdriving module for the steeringWheel")
	s.steering += "Self_drive"
	t.selfdriving = s.steering
}

func (t *smartAuto) visitForGearShift(s *gearshift) {
	fmt.Println("Add automatic shifting for the gearbox")
	s.gear += "Auto_shift"
	t.selfdriving = s.gear
}

type isAuto struct {
	active bool
}

func (a *isAuto) visitForSteeringWheel(s *steeringWheel) {
	fmt.Println("Checking the steeringWheel")
	switch {
	case s.steering == "Self_drive":
		a.active = true
	default:
		a.active = false
	}
}

func (a *isAuto) visitForGearShift(c *gearshift) {
	fmt.Println("Checking the gearshift")
	switch {
	case c.gear == "Auto_shift":
		a.active = true
	default:
		a.active = false
	}
}

func main() {
	st_wheel := &steeringWheel{steering: "Self_drive"}
	gr_shift := &gearshift{gear: "Auto_shift"}

	smart := &smartAuto{}

	st_wheel.accept(smart)
	gr_shift.accept(smart)

	fmt.Println()
	whattype := &isAuto{}
	st_wheel.accept(whattype)
	gr_shift.accept(whattype)
}
