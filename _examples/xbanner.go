package main

import (
	"github.com/robjporter/go-library/xbanner"
)

func main() {
	displaybanner()
}

func displaybanner() {
	xbanner.PrintNewFigure("LIBRARY", "varsity", true)
}
