package main

import (
	"fmt"
	"github.com/robjporter/go-library/xping"
)

func main() {
	pdp := xping.New()
	result := pdp.Ping("10.253.104.97", 5, 1)

	fmt.Println("RESULT: ", result)
}
