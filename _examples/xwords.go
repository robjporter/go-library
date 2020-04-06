package main

import (
	"fmt"
	"math/rand"

	"github.com/robjporter/go-library/xwords"
)

func main() {
	word := xwords.Words[rand.Intn(len(xwords.Words))]
	fmt.Println("Word: ",word)
}
