package pimessage

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

// DisplayConfig Setup configuration for the display hardware.
type DisplayConfig struct {
	// RPi Pin setup
	LatchPin int
	ClockPin int
	DataPin  int

	En74138 int
	La74138 int
	Lb74138 int
	Lc74138 int
	Ld74138 int

	Columns int
	Rows    int
}

// Display An abstraction of a cheap Chinese Dot Matrix display
type Display struct {
	config DisplayConfig
	canvas Canvas
}

// NewDisplay Create and setup a Display
func NewDisplay(conf DisplayConfig) *Display {
	d := new(Display)
	d.config = conf

	// Start the GPIO pins off
	// Get the IO Pins started
	err := rpio.Open()
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	return d
}

// Start Start the display rendering
func (d *Display) Start() error {
	fmt.Println("Display Start")

	return nil
}
