package main

import (
	"fmt"

	"github.com/robjporter/go-library/xeach"
)

func main() {
	fmt.Println("")
	fmt.Println("EACH *******************************************************")
	fn := func(s, i interface{}) {
		fmt.Println(s.(string))
	}
	s := []string{"a", "b", "c"}

	xeach.Each(s, fn)
}