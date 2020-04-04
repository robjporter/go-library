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
	g.Borders.SetBorderStyle("circles")
	g.Borders.SetContent(lines)
	g.Borders.SetSpacer(12)
	g.Borders.PrintBorder(true)
}
