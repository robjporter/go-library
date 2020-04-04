// LOG=* go run logger.go

package main

import (
	"github.com/robjporter/go-library/xlogger"
	"errors"
	"fmt"
	"time"
)

var log = xlogger.New("app")
var app = xlogger.New("app")
var images = xlogger.New("images")
var socket = xlogger.New("websocket")
var users = xlogger.New("users")
var db = xlogger.New("database")

type CustomWriter struct{}

func (customWriter *CustomWriter) Init() {

}

func (cw CustomWriter) Write(log *xlogger.Log) {
	fmt.Println("custom log -> ", log.Package, log.Level, log.Message, log.Attrs)
}

func main() {
// Simple
	log.Info("Starting at %d", 9088)

	log.Info("Requesting an image at foo/bar.jpg")
	timer := log.Timer()
	time.Sleep(time.Millisecond * 250)
	timer.End("Fetched foo/bar.jpg")

	log.Error("Failed to start, shutting down...")

// Timers
	var perf = xlogger.New("performance")
	var test = xlogger.New("test")

	timer = perf.Timer()
	for i := 0; i < 50; i++ {
		t := test.Timer()
		t.End("foobar %s", "yoyo", xlogger.Attrs{
			"foo": 123,
			"bar": true,
		})
	}
	timer.End("End")

// Attributes
	log.Info("Sending an e-mail", xlogger.Attrs{
		"from": "foo@bar.com",
		"to": "qux@corge.com",
	})

	err := errors.New("Too busy")

	log.Error("Failed to send e-mail. Error: %s", err, xlogger.Attrs{
		"from": "foo@bar.com",
		"to": "qux@corge.com",
	})

	timer = log.Timer()
	time.Sleep(time.Millisecond * 500)
	timer.End("Created a new %s image", "bike", xlogger.Attrs{
		"id": 123456,
		"model": "bmx",
		"frame": "purple",
		"year": 2014,
	})

// Custom Writer
	xlogger.Hook(&CustomWriter{})
	log.Info("he-yo")

	log.Info("Requesting an image", xlogger.Attrs{
		"file": "foo/bar.jpg",
	})

	timer = log.Timer()
	time.Sleep(time.Millisecond * 250)
	timer.End("Fetched foo/bar.jpg")

	log.Error("Failed, shutting down...")

// Multi logger
	go app.Info("Starting at %d", 9088)

	go db.Info("Connecting to mysql://azer@localhost:9900/foobar")
	go images.Info("Requesting an image at foo/bar.jpg")

	timer = images.Timer()
	time.Sleep(time.Millisecond * 250)
	timer.End("Fetched foo/bar.jpg")

	db.Error("Fatal connection error.")

	users.Info("%s just logged  from %s", "John", "Istanbul")

	socket.Info("Connecting...")

	err = errors.New("Unable to connect.")
	socket.Error("%v", err)

	time.Sleep(time.Millisecond * 250)

	app.Error("Failed to start, shutting down...")
}
