package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	. "github.com/robjporter/go-library/debug"
)

type MyData struct {
	IntField   int
	FloatField float64
	StrField   string
	MapField   map[int]string
	SliceField []int
	PointField *MyData
}

func main() {
	data := &MyData{
		1234,
		77.88,
		"xyz",
		map[int]string{
			1: "abc",
			2: "def",
			3: "ghi",
		},
		[]int{
			3,
			7,
			11,
			13,
			17,
		},
		nil,
	}
	data.PointField = data

	Print("This is a test")
	fmt.Println(GODEBUG("DISPLAY", "Does Not Exist"))
	fmt.Println(GODEBUG("PWD", "Does Not Exist"))
	fmt.Println(os.Environ())
	fmt.Println(GODEBUGDATA())

	fmt.Println(string(Dump(DumpStyle{Format: true, Indent: "  "}, data)))
	fmt.Printf("\n%s", Dump(DumpStyle{Pointer: true, Indent: "    "}, data))
	fmt.Printf("\n%s", Dump(DumpStyle{Format: true, Indent: "    "}, data))
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