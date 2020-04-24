package main

import (
	"fmt"
	"time"

	"github.com/robjporter/go-library/xtimeout"
)

func main() {
	a := xtimeout.WithTimeout(test, time.Second * 2)
	fmt.Println("Feedback: ", a)
}

func test() {
	time.Sleep(time.Second * 5)
}
