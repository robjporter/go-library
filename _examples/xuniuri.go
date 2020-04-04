package main

import (
	"fmt"
	"github.com/robjporter/go-library/xuniuri"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,/-=[]{};:~")

func main() {
	s := xuniuri.New()
	fmt.Println("Default URI - New: ", s)

	s = xuniuri.NewLen(40)
	fmt.Println("Default URI - NewLen: ", s)

	s = xuniuri.NewLenChars(40, StdChars)
	fmt.Println("Default URI - NewLenChars: ", s)
}
