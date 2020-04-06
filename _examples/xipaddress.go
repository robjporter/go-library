package main

import (
	"fmt"
	"net"
	"github.com/robjporter/go-library/xipaddress"
)

func main() {
	ip := net.ParseIP("1.2.3.4")

	ipn, err := xipaddress.IpToUint32(ip)
	if err != nil {
		fmt.Printf("Error")
	}

	fmt.Printf("ip %s is %d as an integer\n", ip, ipn)

	ips := xipaddress.Uint32ToIP(ipn).String()

	fmt.Printf("And we used Uint32ToIP to return it to %s\n", ips)

	_, netw, _ := net.ParseCIDR("1.2.3.4/24")
	last := xipaddress.LastAddress(netw)

	fmt.Printf("The last address in the net range is %s\n", last.String())

	//LastAddress works for ipv6 too
	_, net6, _ := net.ParseCIDR("2001:658:22a:cafe::/64")
	last = xipaddress.LastAddress(net6)

	fmt.Printf("The last address in the ipv6 network range is %s\n", last.String())
}
