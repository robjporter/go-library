package main

import (
	"fmt"
	"time"

	"../xspinners"
)

func main() {
	spin := xspinners.New(xspinners.DOTS1, " TEST")
	fmt.Println(spin)
	spin.Start()
	time.Sleep(5 * time.Second)
	spin.Stop()

	spin = xspinners.New(xspinners.DOTS2, " TEST")
	fmt.Println(spin)
	spin.Start()
	time.Sleep(5 * time.Second)
	spin.Stop()

	spin = xspinners.New(xspinners.CIRCLE1, " TEST")
	fmt.Println(spin)
	spin.Start()
	time.Sleep(5 * time.Second)
	spin.Stop()

	spin = xspinners.New(xspinners.ARROWS1, " TEST")
	fmt.Println(spin)
	spin.Start()
	time.Sleep(5 * time.Second)
	spin.Stop()
}
