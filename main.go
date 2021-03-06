package main

import (
	"fmt"

	"github.com/rtbaker/GoPiMessage/pimessage"
)

func main() {
	//var mine pimessage.Font;
	//mine = pimessage.Minimal_6x5

	//fmt.Println("hello", mine)

	// Set up the display
	displayConf := pimessage.DisplayConfig{
		LatchPin: 27,
		ClockPin: 28,
		DataPin:  29,
		En74138:  21,
		La74138:  22,
		Lb74138:  23,
		Lc74138:  24,
		Ld74138:  25,
		Columns:  64,
		Rows:     16,
	}

	display := pimessage.NewDisplay(displayConf)

	if display == nil {
		fmt.Println("Problem starting display")
		return
	}

	// Make sure we clean up at the end
	defer display.Finish()

	error := display.Start()
	if error != nil {
		fmt.Println("Error starting display: ", error)
	}
}
