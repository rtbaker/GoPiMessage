package pimessage

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// DisplayConfig Setup configuration for the display hardware.
type DisplayConfig struct {
	// RPi Pin setup
	LatchPin uint32
	ClockPin uint32
	DataPin  uint32

	En74138 uint32
	La74138 uint32
	Lb74138 uint32
	Lc74138 uint32
	Ld74138 uint32

	Columns uint32
	Rows    uint32
}

// Display An abstraction of a cheap Chinese Dot Matrix display
type Display struct {
	config DisplayConfig
	canvas Canvas

	// Private stuff (pins etc)
	latchPin   rpio.Pin
	clockPin   rpio.Pin
	dataPin    rpio.Pin
	en74138Pin rpio.Pin
	la74138Pin rpio.Pin
	lb74138Pin rpio.Pin
	lc74138Pin rpio.Pin
	ld74138Pin rpio.Pin
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

	// Set the pin modes
	d.latchPin = rpio.Pin(conf.LatchPin)
	d.latchPin.Output()

	d.clockPin = rpio.Pin(conf.ClockPin)
	d.clockPin.Output()

	d.dataPin = rpio.Pin(conf.DataPin)
	d.dataPin.Output()

	d.en74138Pin = rpio.Pin(conf.En74138)
	d.en74138Pin.Output()

	d.la74138Pin = rpio.Pin(conf.La74138)
	d.la74138Pin.Output()

	d.lb74138Pin = rpio.Pin(conf.Lb74138)
	d.lb74138Pin.Output()

	d.lc74138Pin = rpio.Pin(conf.Lc74138)
	d.lc74138Pin.Output()

	d.ld74138Pin = rpio.Pin(conf.Ld74138)
	d.ld74138Pin.Output()

	d.en74138Pin.Low()

	return d
}

// Start Start the display rendering
func (d *Display) Start() error {
	fmt.Println("Display Start called")

	return nil
}

// Finish Call before programme end
func (d *Display) Finish() {
	fmt.Println("Display Finish called")
	rpio.Close()
}

/* ----------------------------------------------------------------------------------------------------------------------------
  Private Functions
*/

func (d *Display) shiftOut(dataOut byte) {
	sleepDuration := 5 * time.Microsecond

	var i uint32
	for i = 0; i <= 7; i++ {
		d.clockPin.Low()

		time.Sleep(sleepDuration)

		if (dataOut & (0x01 << i)) > 0 {
			d.dataPin.High()
		} else {
			d.dataPin.Low()
		}

		time.Sleep(sleepDuration)

		d.clockPin.High()

		time.Sleep(sleepDuration)
	}
}

func (d *Display) displayMatrix(BMP []uint8) {
	//Display count
	var disCnt uint32 = 4 //256;
	var i uint32
	var RowXPixel uint32

	for i = 0; i < disCnt*16; i++ {
		// What to display
		ColNum1 := ^BMP[(d.config.Columns/8)*RowXPixel]
		ColNum2 := ^BMP[(d.config.Columns/8)*RowXPixel+1]
		ColNum3 := ^BMP[(d.config.Columns/8)*RowXPixel+2]
		ColNum4 := ^BMP[(d.config.Columns/8)*RowXPixel+3]
		ColNum5 := ^BMP[(d.config.Columns/8)*RowXPixel+4]
		ColNum6 := ^BMP[(d.config.Columns/8)*RowXPixel+5]
		ColNum7 := ^BMP[(d.config.Columns/8)*RowXPixel+6]
		ColNum8 := ^BMP[(d.config.Columns/8)*RowXPixel+7]

		// Display off
		//      PIN_MAP[en_74138].gpio_peripheral->BSRR = PIN_MAP[en_74138].gpio_pin; //digitalWrite(en_74138, HIGH);

		//Col scanning
		d.shiftOut(ColNum1)
		d.shiftOut(ColNum2)
		d.shiftOut(ColNum3)
		d.shiftOut(ColNum4)
		d.shiftOut(ColNum5)
		d.shiftOut(ColNum6)
		d.shiftOut(ColNum7)
		d.shiftOut(ColNum8)

		// toggle the latch
		d.latchPin.Low()
		//      delayMicroseconds(20);
		d.latchPin.High()

		//Row scanning
		// AVR Port Operation
		// PORTD = ((RowXPixel << 3 ) & 0X78) | (PORTD & 0X87);//Write PIN  la_74138 lb_74138 lc_74138 ld_74138
		if (RowXPixel & 0x1) > 0 {
			d.la74138Pin.High()
		} else {
			d.la74138Pin.Low()
		}

		if (RowXPixel & 0x2) > 0 {
			d.lb74138Pin.High()
		} else {
			d.lb74138Pin.Low()
		}

		if (RowXPixel & 0x4) > 0 {
			d.lc74138Pin.High()
		} else {
			d.lc74138Pin.Low()
		}

		if (RowXPixel & 0x8) > 0 {
			d.ld74138Pin.High()
		} else {
			d.ld74138Pin.Low()
		}

		// Display on
		//      PIN_MAP[en_74138].gpio_peripheral->BRR = PIN_MAP[en_74138].gpio_pin; //digitalWrite(en_74138, LOW);
		// Next row
		if RowXPixel == 15 {
			RowXPixel = 0
		} else {
			RowXPixel++
		}

		// delayMicroseconds(1000);
	}

}
