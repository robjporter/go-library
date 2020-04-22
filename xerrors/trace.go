package xerrors

import "runtime"

func trace() StackTrace {
	// Retrieve up to 10 PC call frames (which can expand to a large number of real Frames)
	pc := make([]uintptr, 10)
	// 0: Callers
	// 1: trace
	// 2: New, Errorf, etc
	// 3: first one we care about
	n := runtime.Callers(3, pc)
	if n == 0 {
		return nil
	}
	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	ci := CallersFrames(pc)
	var st StackTrace
	for {
		frame, more := ci.Next()
		// `frame` is valid even if `more` is false, so append it here
		f := new(runtime.Frame)
		*f = frame
		st = append(st, f)
		if !more {
			break
		}
	}
	return st
}

/*
	Note that github.com/pkg/errors misuses the golang runtime API.  The correct method of retrieving a stack trace
	was exemplified [here](https://golang.org/pkg/runtime/#Frames):
```go
package main
import (
	"fmt"
	"runtime"
	"strings"
)
func main() {
	c := func() {
		// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			// No pcs available. Stop now.
			// This can happen if the first argument to runtime.Callers is large.
			return
		}
		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
		frames := runtime.CallersFrames(pc)
		// Loop to get frames.
		// A fixed number of pcs can expand to an indefinite number of Frames.
		for {
			frame, more := frames.Next()
			// To keep this example's output stable
			// even if there are changes in the testing package,
			// stop unwinding when we leave package runtime.
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)
			if !more {
				break
			}
		}
	}
	b := func() { c() }
	a := func() { b() }
	a()
}
```
*/
