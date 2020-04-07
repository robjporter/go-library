package main

import (
	"fmt"
	"time"

	"../xtimer"
)

func main() {

	xtimer.Timer("main")

	time.Sleep(5 * time.Second)

	xtimer.Timer("step 2")
	time.Sleep(2 * time.Second)
	fmt.Println(xtimer.Timer("step 2"))

	fmt.Println(xtimer.Timer("main"))
}
