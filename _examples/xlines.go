package main

import (
	"fmt"
	"strconv"

	"github.com/robjporter/go-library/xlines"
)

func main() {
	g := xlines.New()
	g.SetLineStyle("line0")
	g.PrintLine(60)
	g.PrintLineTitleLeft(60, "Results")
	g.PrintLineTitleCenter(60, "* Results *")
	g.PrintLineTitleRight(60, "Results")

	for i := 0; i < 13; i++ {
		fmt.Println("Style line" + strconv.Itoa(i) )
		g.SetLineStyle("line" + strconv.Itoa(i))
		g.PrintLine(60)
	}
}
