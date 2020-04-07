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
	fmt.Println("Is Timer 1 Active: ", stopWatch.IsActive("[<step 1>]"))
	fmt.Println("Active Timers Count: ", stopWatch.ActiveTimers())
	stopWatch.Stop("[<step 1>]")

	fmt.Println("Is Timer 1 Active: ", stopWatch.IsActive("[<step 1>]"))
	fmt.Println("Active Timers Count: ", stopWatch.ActiveTimers())
	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 2>]")
	time.Sleep(time.Millisecond * 2)
	stopWatch.Stop("[<step 2>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 3>]")
	time.Sleep(time.Millisecond * 3)
	stopWatch.Stop("[<step 3>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 4>]")
	time.Sleep(time.Millisecond * 4)
	stopWatch.Stop("[<step 4>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 5>]")
	time.Sleep(time.Millisecond * 100)
	stopWatch.Stop("[<step 5>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 6>]")
	time.Sleep(time.Millisecond * 2)
	stopWatch.Stop("[<step 6>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 7>]")
	time.Sleep(time.Millisecond * 3)
	stopWatch.Stop("[<step 7>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	stopWatch.Start("[<step 8>]")
	time.Sleep(time.Millisecond * 4)
	stopWatch.Stop("[<step 8>]")

	fmt.Println("Last Timer Started: ", stopWatch.LastTaskStarted())

	fmt.Println(stopWatch.PrettyPrint())
}
