package main

import (
	"github.com/robjporter/go-library/xicons"
)

func main() {
	g := xicons.New()
	g.PrintIconStyles()
	g.PrintIcon("cross")
	g.PrintIcon("alien")
}