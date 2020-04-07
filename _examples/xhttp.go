package main

import (
	"fmt"

	"../xhttp"
	"time"
)

func main() {
	s := xhttp.NewSite()
	s.SetURL("http://www.google.co.uk")
	fmt.Println(s.Status())
	fmt.Println(s.LastStatus())
	fmt.Println(s.LastStatusString())

	// TODO:
	h := xhttp.NewHammer()
	h.SetThreadCount(4)
	h.SetHitCount(20)
	h.SetDelay(100*time.Millisecond)
	h.SetSpread(1*time.Second)
	h.SetURL("http://localhost:8080")
	h.Run()

	fmt.Println("HITS:     ",h.GetHits())
	fmt.Println("DURATION: ",h.GetDuration())
	fmt.Println("FAILED:   ",h.GetFailed())
	fmt.Println("QUICKEST: ",h.GetQuickest())
	fmt.Println("SLOWEST:  ",h.GetSlowest())

	res, err := xhttp.GetHTTP("https://jsonplaceholder.typicode.com/todos/4", []string{})

	fmt.Println("GET ERROR: ", err)
	fmt.Println("GET RESULTS: ", res)

	resp, err := xhttp.PostHTTP("https://jsonplaceholder.typicode.com/posts","{title: 'foo',body: 'bar',userId: 1}",[]string{"Content-type=application/json; charset=UTF-8"})
	fmt.Println("POST ERROR: ", err)
	fmt.Println("POST RESULTS: ", resp)

}
