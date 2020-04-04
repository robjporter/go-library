package main

import (
	"github.com/robjporter/go-library/xlines"
)

func main() {
	g := xlines.New()
	g.SetLineStyle("line0")
	g.PrintLine(60)
	g.PrintLineTitleLeft(60, "Results")
	g.PrintLineTitleCenter(60, "* Results *")
	g.PrintLineTitleRight(60, "Results")
}
