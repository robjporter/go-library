package main

import (
	"fmt"
	"github.com/robjporter/go-library/xtreemap"
)


func less(x, y int) bool { return x < y }

func main() {
	tr := xtreemap.New(less)
	tr.Set(0, "Hello")
	tr.Set(1, "World")

	for it := tr.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}