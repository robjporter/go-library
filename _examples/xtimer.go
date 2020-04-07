package main

import (
	"fmt"
	"time"

	"github.com/robjporter/go-library/xtimer"
)

func main() {
	stopWatch := xtimer.New("Testing")

	stopWatch.Start("[<step 1>]")
	time.Sleep(time.Millisecond * 100)
	stopWatch.Stop("[<step 1>]")

	stopWatch.Start("[<step 2>]")
	time.Sleep(time.Millisecond * 2)
	stopWatch.Stop("[<step 2>]")

	stopWatch.Start("[<step 3>]")
	time.Sleep(time.Millisecond * 3)
	stopWatch.Stop("[<step 3>]")

	stopWatch.Start("[<step 4>]")
	time.Sleep(time.Millisecond * 4)
	stopWatch.Stop("[<step 4>]")

	fmt.Println(stopWatch.PrettyPrint())
}
