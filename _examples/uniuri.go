package main

import (
	"fmt"
	"github.com/robjporter/go-library/uniuri"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,/-=[]{};:~")

func main() {
	s := uniuri.New()
	fmt.Println("Default URI - New: ", s)

	s = uniuri.NewLen(40)
	fmt.Println("Default URI - NewLen: ", s)

	s = uniuri.NewLenChars(40, StdChars)
	fmt.Println("Default URI - NewLenChars: ", s)
}
