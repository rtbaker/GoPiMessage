package main

import (
	"fmt"
	"github.com/rtbaker/GoPiMessage/pimessage"
)

func main() {
	var mine pimessage.Font;
	mine = pimessage.Minimal_6x5

	fmt.Println("hello", mine)

}
