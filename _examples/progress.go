package main

import (
	"math/rand"
	"time"

	"github.com/robjporter/go-library/mpb"
	"github.com/robjporter/go-library/mpb/decor"
)

func main() {
	// initialize progress container, with custom width
	p := multiprogress.New(multiprogress.WithWidth(64))

	total := 100
	name := "Single Bar:"
	// adding a single bar, which will inherit container's width
	bar := p.AddBar(int64(total),
		// override DefaultBarStyle, which is "[=>-]<+"
		multiprogress.BarStyle("╢▌▌░╟"),
		multiprogress.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		multiprogress.AppendDecorators(decor.Percentage()),
	)
	// simulating some work
	max := 100 * time.Millisecond
	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		bar.Increment()
	}
	// wait for our bar to complete and flush
	p.Wait()
}
