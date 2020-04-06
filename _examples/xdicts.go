package main

import (
	"fmt"

	"github.com/robjporter/go-library/xcounter"
	"github.com/robjporter/go-library/xdefaultdict"
	"github.com/robjporter/go-library/xordereddict"
)

func main() {
	d := xordereddict.New()
	d.Set("foo", "bar")
	d.Set("baz", 123)

	for v := range d.Iterate() {
		fmt.Println(v)
	}

	c := xcounter.New()
	c.AddItems("foo", "foo", "bar", "baz")

	for i, item := range c.MostCommon(5) {
		fmt.Printf("%d: %#v seen %d time(s)\n", i+1, item.Value, item.Count)
	}

	dict := xdefaultdict.New(xdefaultdict.IntDefault)
	fmt.Printf("foo default: %d\n", dict.Get("foo"))
}
