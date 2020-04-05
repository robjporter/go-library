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
	fmt.Println("EACH *******************************************************")
	xeach.Each(s, fn)
}