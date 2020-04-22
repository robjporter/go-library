package main

import (
	"fmt"
	"github.com/robjporter/go-library/xemoji"
)

func main() {
	a := xemoji.GetEmoji(0)
	b := xemoji.GetEmojis([]int{0,1,2,3,4})

	fmt.Println(a)
	fmt.Println(b)

	xemoji.OutputAll()
}
