package main

import (
	"fmt"

	"github.com/robjporter/go-library/xdisplay"
)

func main() {
	d := xdisplay.New()
	d.ClearScreen()

	fmt.Println("Terminal Size: ", d.TerminalSize())
}
