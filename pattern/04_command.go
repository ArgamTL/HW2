package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type headlight struct {
	isRunning bool
}

func (hl *headlight) on() {
	hl.isRunning = true
	fmt.Println("Turning headlights on")
}

func (hl *headlight) off() {
	hl.isRunning = false
	fmt.Println("Turning headlights off")
}

func main() {
	headlight := &headlight{}
	onCommand := &onCommand{
		device: headlight,
	}
	offCommand := &offCommand{
		device: headlight,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
