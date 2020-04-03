package main

import (
	"fmt"
	"github.com/robjporter/go-library/ipify"
)

func main() {
	ip, err := ipify.GetIp()
	if err != nil {
		fmt.Println("Couldn't get my IP address:", err)
	} else {
		fmt.Println("My IP address is:", ip)
	}
}
