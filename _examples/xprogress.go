package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/robjporter/go-library/xmpb"
	"github.com/robjporter/go-library/xmpb/decor"
)

func main() {
	// initialize progress container, with custom width
	p := xmpb.New(xmpb.WithWidth(64))

	total := 100
	name := "Single Bar:"
	// adding a single bar, which will inherit container's width
	bar := p.AddBar(int64(total),
		// override DefaultBarStyle, which is "[=>-]<+"
		xmpb.BarStyle("╢▌▌░╟"),
		xmpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		xmpb.AppendDecorators(decor.Percentage()),
	)
	// simulating some work
	max := 100 * time.Millisecond
	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		bar.Increment()
	}
	// wait for our bar to complete and flush
	p.Wait()

	var wg sync.WaitGroup
	// pass &wg (optional), so p will wait for it eventually
	p = xmpb.New(xmpb.WithWaitGroup(&wg))
	total, numBars := 100, 3
	wg.Add(numBars)

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Bar#%d:", i)
		bar := p.AddBar(int64(total),
			xmpb.PrependDecorators(
				// simple name decorator
				decor.Name(name),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
			xmpb.AppendDecorators(
				// replace ETA decorator with "done" message, OnComplete event
				decor.OnComplete(
					// ETA decorator with ewma age of 60
					decor.EwmaETA(decor.ET_STYLE_GO, 60), "done",
				),
			),
		)
		// simulating some work
		go func() {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			max := 100 * time.Millisecond
			for i := 0; i < total; i++ {
				// start variable is solely for EWMA calculation
				// EWMA's unit of measure is an iteration's duration
				start := time.Now()
				time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
				bar.Increment()
				// we need to call DecoratorEwmaUpdate to fulfill ewma decorator's contract
				bar.DecoratorEwmaUpdate(time.Since(start))
			}
		}()
	}
	// Waiting for passed &wg and for all bars to complete and flush
	p.Wait()
}
