package main

import (
	"fmt"
	"github.com/robjporter/go-library/xping"
)

func main() {
	pdp := xping.New()
	result := pdp.Ping("10.237.104.97", 5, 1)

	fmt.Println("RESULT: ", result)
	fmt.Println("RESULT ONLINE: ", result.Online)
	fmt.Println("RESULT PACKETLOSS: ", result.PacketLoss)
	fmt.Println("RESULT LATENCY: ", result.Latency)

}
