package main

import (
	"fmt"

	"github.com/robjporter/go-library/xdisplay"
)

func main() {
	d := xdisplay.New()
	d.ClearScreen()

	a, b, err := d.TerminalSize()
	fmt.Println("Terminal Size: ", a, "x", b, "x", err)
}
