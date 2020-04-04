package main

import (
	"github.com/robjporter/go-library/xborders"
)

func main() {
	var lines []string
	lines = append(lines, "")
	lines = append(lines, "This is the title")
	lines = append(lines, "")

	g := xborders.New()
	g.SetBorderStyle("circles")
	g.SetContent(lines)
	g.SetSpacer(12)
	g.PrintBorder(true)

	g.SetBorderStyle("say")
	g.PrintBorder(true)

	g.SetBorderStyle("classicish")
	g.PrintBorder(true)

	g.SetBorderStyle("think")
	g.PrintBorder(true)

	g.SetBorderStyle("unicode")
	g.PrintBorder(true)

	g.SetBorderStyle("thick")
	g.PrintBorder(true)

	g.SetBorderStyle("rounded")
	g.PrintBorder(true)

	g.SetBorderStyle("bigger")
	g.PrintBorder(true)

	g.SetBorderStyle("fancy")
	g.PrintBorder(true)

	g.SetBorderStyle("fancy2")
	g.PrintBorder(true)

	g.SetBorderStyle("blocks")
	g.PrintBorder(true)
}
