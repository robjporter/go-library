package main

import (
	"fmt"

	"github.com/robjporter/go-library/xarchive"
)

func main() {
	e := xarchive.Zipit("./", "archive.zip")
	fmt.Println("ERROR: ",e)
	e = xarchive.Unzipit("./archive.zip", "archive")
	fmt.Println("ERROR: ",e)
}
