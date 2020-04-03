package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"github.com/robjporter/go-library/debug"
)

func main() {
	Print("This is a test")
	fmt.Println(GODEBUG(""))
	fmt.Println(GODEBUGDATA)
}

func Print(v ...interface{}) {
	fmt.Fprintf(os.Stderr, "[DEBUG PRINT]\nby goroutine %s\n%s\n", GoroutineID(), StackTrace(2).Bytes(""))
	fmt.Fprintf(os.Stderr, string(Dump(DumpStyle{Format: true, Indent: ""}, v...)))
}

func GoroutineID() string {
	buf := make([]byte, 20)
	buf = buf[:runtime.Stack(buf, false)]
	return string(bytes.Split(buf, []byte(" "))[1])
}