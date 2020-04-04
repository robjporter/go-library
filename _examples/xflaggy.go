package main

import (
	"fmt"
	"github.com/robjporter/go-library/xflaggy"
)

func main() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"
	var stringFlag2 = "defaultValue"

	// Add a flag
	xflaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Create the subcommand
	subcommand := xflaggy.NewSubcommand("subcommandExample")

	// Add a flag to the subcommand
	subcommand.String(&stringFlag2, "g", "flag", "A test string flag")

	// Add the subcommand to the parser at position 1
	xflaggy.AttachSubcommand(subcommand, 1)

	// Parse the flag
	xflaggy.Parse()

	// Use the flag
	fmt.Println("stringFlag: ",stringFlag)
	fmt.Println("stringFlag2: ",stringFlag2)
}
