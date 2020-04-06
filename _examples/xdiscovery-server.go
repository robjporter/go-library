package main

import (
	"fmt"
	"time"

	"github.com/robjporter/go-library/xdiscovery"
)

func main() {
	port := "1234"
	auth := xdiscovery.NewBasicAuthenticator("roporter", "password")
	server := xdiscovery.NewRandomServer(port, auth)
	err := server.ListenAndServe()

	fmt.Println("STARTING SERVER: ", err)

	discoverHost := "localhost:1234"
	timeOut, _ := time.ParseDuration("30s")
	client1, err1 := xdiscovery.NewRegistryClient("client1", "", discoverHost, auth, timeOut)
	client2, err2 := xdiscovery.NewRegistryClient("client2", "", discoverHost, auth, timeOut)

	fmt.Println("CLIENT 1 CREATION ERROR: ", err1)
	fmt.Println("CLIENT 2 CREATION ERROR: ", err2)

	renew, _ := time.ParseDuration("15s")
	err1 = client1.Auto(renew)
	err2 = client2.Auto(renew)

	fmt.Println("CLIENT 1 REGISTER ERROR: ", err1)
	fmt.Println("CLIENT 2 REGISTER ERROR: ", err2)

	services1, err1 := client1.List("client1")
	services2, err2 := client1.List("client2")

	fmt.Println("CLIENT 1 SERVICES ERROR: ", err1)
	fmt.Println("CLIENT 2 SERVICES ERROR: ", err2)

	fmt.Println("CLIENT 1 SERVICES: ", services1)
	fmt.Println("CLIENT 2 SERVICES: ", services2)

	host1, err1 := client1.Discover("client2")
	host2, err2 := client2.Discover("client1")

	fmt.Println("CLIENT 1 DISCOVER ERROR: ", err1)
	fmt.Println("CLIENT 2 DISCOVER ERROR: ", err2)

	fmt.Println("CLIENT 1 DISCOVER: ", host1)
	fmt.Println("CLIENT 2 DISCOVER: ", host2)

	err1 = client1.Deregister()
	err2 = client2.Deregister()

	fmt.Println("CLIENT 1 DEREGISTER ERROR: ", err1)
	fmt.Println("CLIENT 2 DEREGISTER ERROR: ", err2)

}
