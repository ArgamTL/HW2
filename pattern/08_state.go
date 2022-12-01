package main

import (
	"fmt"
)

// PC state on|off
type PC struct {
	current State
}

// NewPC
func NewPC() *PC {
	fmt.Printf("PC is ready.\n")
	return &PC{NewOFF()}
}

// setCurrent
func (m *PC) setCurrent(s State) {
	m.current = s
}

// On
func (m *PC) On() {
	m.current.On(m)
}

// Off
func (m *PC) Off() {
	m.current.Off(m)
}

// State
type State interface {
	On(m *PC)
	Off(m *PC)
}

// ON
type ON struct {
}

func NewON() State {
	return &ON{}
}

// --
func (o *ON) On(m *PC) {
	fmt.Printf("   already ON\n")
}

// Off
func (o *ON) Off(m *PC) {
	fmt.Printf("going from ON to OFF\n")
	m.setCurrent(NewOFF())
}

// OFF
type OFF struct {
}

func NewOFF() State {
	return &OFF{}
}

func (o *OFF) On(m *PC) {
	fmt.Printf("going from OFF to ON\n")
	m.setCurrent(NewON())
}

// --
func (o *OFF) Off(m *PC) {
	fmt.Printf("already OFF\n")
}

func main() {

	comp := NewPC()
	comp.Off()
	comp.On()
	comp.On()
	comp.Off()

}
