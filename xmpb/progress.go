package xmpb

import (
	"bytes"
	"container/heap"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/robjporter/go-library/xmpb/cwriter"
	"github.com/robjporter/go-library/xmpb/decor"
)

const (
	// default RefreshRate
	prr = 120 * time.Millisecond
	// default width
	pwidth = 80
)

// Progress represents the container that renders Progress bars
type Progress struct {
	ctx          context.Context
	uwg          *sync.WaitGroup
	cwg          *sync.WaitGroup
	bwg          *sync.WaitGroup
	operateState chan func(*pState)
	done         chan struct{}
	refreshCh    chan time.Time
	once         sync.Once
	dlogger      *log.Logger
}

type pState struct {
	bHeap            priorityQueue
	heapUpdated      bool
	pMatrix          map[int][]chan int
	aMatrix          map[int][]chan int
	barShutdownQueue []*Bar
	barPopQueue      []*Bar

	// following are provided/overrided by user
	idCount          int
	width            int
	popCompleted     bool
	rr               time.Duration
	uwg              *sync.WaitGroup
	refreshSrc       <-chan time.Time
	renderDelay      <-chan struct{}
	shutdownNotifier chan struct{}
	parkedBars       map[*Bar]*Bar
	output           io.Writer
	debugOut         io.Writer
}

// New creates new Progress container instance. It's not possible to
// reuse instance after *Progress.Wait() method has been called.
func New(options ...ContainerOption) *Progress {
	return NewWithContext(context.Background(), options...)
}

// NewWithContext creates new Progress container instance with provided
// context. It's not possible to reuse instance after *Progress.Wait()
// method has been called.
func NewWithContext(ctx context.Context, options ...ContainerOption) *Progress {
	s := &pState{
		bHeap:      priorityQueue{},
		width:      pwidth,
		rr:         prr,
		parkedBars: make(map[*Bar]*Bar),
		output:     os.Stdout,
		debugOut:   ioutil.Discard,
	}

	for _, opt := range options {
		if opt != nil {
			opt(s)
		}
	}

	p := &Progress{
		ctx:          ctx,
		uwg:          s.uwg,
		cwg:          new(sync.WaitGroup),
		bwg:          new(sync.WaitGroup),
		operateState: make(chan func(*pState)),
		done:         make(chan struct{}),
		dlogger:      log.New(s.debugOut, "[mpb] ", log.Lshortfile),
	}

	p.cwg.Add(1)
	go p.serve(s, cwriter.New(s.output))
	return p
}

// AddBar creates a new progress bar and adds it to the rendering queue.
func (p *Progress) AddBar(total int64, options ...BarOption) *Bar {
	return p.Add(total, NewBarFiller(DefaultBarStyle, false), options...)
}

// AddSpinner creates a new spinner bar and adds it to the rendering queue.
func (p *Progress) AddSpinner(total int64, alignment SpinnerAlignment, options ...BarOption) *Bar {
	return p.Add(total, NewSpinnerFiller(DefaultSpinnerStyle, alignment), options...)
}

// Add creates a bar which renders itself by provided filler.
// Set total to 0, if you plan to update it later.
// Panics if *Progress instance is done, i.e. called after *Progress.Wait().
func (p *Progress) Add(total int64, filler BarFiller, options ...BarOption) *Bar {
	if filler == nil {
		filler = NewBarFiller(DefaultBarStyle, false)
	}
	p.bwg.Add(1)
	result := make(chan *Bar)
	select {
	case p.operateState <- func(ps *pState) {
		bs := ps.makeBarState(total, filler, options...)
		bar := newBar(p, bs)
		if bs.runningBar != nil {
			bs.runningBar.noPop = true
			ps.parkedBars[bs.runningBar] = bar
		} else {
			heap.Push(&ps.bHeap, bar)
			ps.heapUpdated = true
		}
		ps.idCount++
		result <- bar
	}:
		bar := <-result
		bar.subscribeDecorators()
		return bar
	case <-p.done:
		p.bwg.Done()
		panic(fmt.Sprintf("%T instance can't be reused after it's done!", p))
	}
}

func (p *Progress) dropBar(b *Bar) {
	select {
	case p.operateState <- func(s *pState) {
		if b.index < 0 {
			return
		}
		heap.Remove(&s.bHeap, b.index)
		s.heapUpdated = true
	}:
	case <-p.done:
	}
}

func (p *Progress) setBarPriority(b *Bar, priority int) {
	select {
	case p.operateState <- func(s *pState) {
		if b.index < 0 {
			return
		}
		b.priority = priority
		heap.Fix(&s.bHeap, b.index)
	}:
	case <-p.done:
	}
}

// UpdateBarPriority same as *Bar.SetPriority(int).
func (p *Progress) UpdateBarPriority(b *Bar, priority int) {
	p.setBarPriority(b, priority)
}

// BarCount returns bars count
func (p *Progress) BarCount() int {
	result := make(chan int, 1)
	select {
	case p.operateState <- func(s *pState) { result <- s.bHeap.Len() }:
		return <-result
	case <-p.done:
		return 0
	}
}

// Wait waits far all bars to complete and finally shutdowns container.
// After this method has been called, there is no way to reuse *Progress
// instance.
func (p *Progress) Wait() {
	if p.uwg != nil {
		// wait for user wg
		p.uwg.Wait()
	}

	// wait for bars to quit, if any
	p.bwg.Wait()

	p.once.Do(p.shutdown)

	// wait for container to quit
	p.cwg.Wait()
}

func (p *Progress) shutdown() {
	close(p.done)
}

func (p *Progress) serve(s *pState, cw *cwriter.Writer) {
	defer p.cwg.Done()

	p.refreshCh = s.newTicker(p.done)

	for {
		select {
		case op := <-p.operateState:
			op(s)
		case <-p.refreshCh:
			if err := s.render(cw); err != nil {
				go p.dlogger.Println(err)
			}
		case <-s.shutdownNotifier:
			return
		}
	}
}

func (s *pState) render(cw *cwriter.Writer) error {
	if s.heapUpdated {
		s.updateSyncMatrix()
		s.heapUpdated = false
	}
	syncWidth(s.pMatrix)
	syncWidth(s.aMatrix)

	tw, err := cw.GetWidth()
	if err != nil {
		tw = s.width
	}
	for i := 0; i < s.bHeap.Len(); i++ {
		bar := s.bHeap[i]
		go bar.render(tw)
	}

	return s.flush(cw)
}

func (s *pState) flush(cw *cwriter.Writer) error {
	var lineCount int
	bm := make(map[*Bar]struct{}, s.bHeap.Len())
	for s.bHeap.Len() > 0 {
		b := heap.Pop(&s.bHeap).(*Bar)
		cw.ReadFrom(<-b.frameCh)
		if b.toShutdown {
			// shutdown at next flush
			// this ensures no bar ends up with less than 100% rendered
			defer func() {
				s.barShutdownQueue = append(s.barShutdownQueue, b)
			}()
		}
		lineCount += b.extendedLines + 1
		bm[b] = struct{}{}
	}

	for _, b := range s.barShutdownQueue {
		if parkedBar := s.parkedBars[b]; parkedBar != nil {
			parkedBar.priority = b.priority
			heap.Push(&s.bHeap, parkedBar)
			delete(s.parkedBars, b)
			b.toDrop = true
		}
		if b.toDrop {
			delete(bm, b)
			s.heapUpdated = true
		} else if s.popCompleted {
			if b := b; !b.noPop {
				defer func() {
					s.barPopQueue = append(s.barPopQueue, b)
				}()
			}
		}
		b.cancel()
	}
	s.barShutdownQueue = s.barShutdownQueue[0:0]

	for _, b := range s.barPopQueue {
		delete(bm, b)
		s.heapUpdated = true
		lineCount -= b.extendedLines + 1
	}
	s.barPopQueue = s.barPopQueue[0:0]

	for b := range bm {
		heap.Push(&s.bHeap, b)
	}

	return cw.Flush(lineCount)
}

func (s *pState) newTicker(done <-chan struct{}) chan time.Time {
	ch := make(chan time.Time)
	if s.shutdownNotifier == nil {
		s.shutdownNotifier = make(chan struct{})
	}
	go func() {
		if s.renderDelay != nil {
			<-s.renderDelay
		}
		if s.refreshSrc == nil {
			ticker := time.NewTicker(s.rr)
			defer ticker.Stop()
			s.refreshSrc = ticker.C
		}
		for {
			select {
			case tick := <-s.refreshSrc:
				ch <- tick
			case <-done:
				close(s.shutdownNotifier)
				return
			}
		}
	}()
	return ch
}

func (s *pState) updateSyncMatrix() {
	s.pMatrix = make(map[int][]chan int)
	s.aMatrix = make(map[int][]chan int)
	for i := 0; i < s.bHeap.Len(); i++ {
		bar := s.bHeap[i]
		table := bar.wSyncTable()
		pRow, aRow := table[0], table[1]

		for i, ch := range pRow {
			s.pMatrix[i] = append(s.pMatrix[i], ch)
		}

		for i, ch := range aRow {
			s.aMatrix[i] = append(s.aMatrix[i], ch)
		}
	}
}

func (s *pState) makeBarState(total int64, filler BarFiller, options ...BarOption) *bState {
	bs := &bState{
		total:    total,
		baseF:    extractBaseFiller(filler),
		filler:   filler,
		priority: s.idCount,
		id:       s.idCount,
		width:    s.width,
		debugOut: s.debugOut,
		extender: func(r io.Reader, _ int, _ *decor.Statistics) (io.Reader, int) {
			return r, 0
		},
	}

	for _, opt := range options {
		if opt != nil {
			opt(bs)
		}
	}

	if s.popCompleted && !bs.noPop {
		bs.priority = -1
	}

	bs.bufP = bytes.NewBuffer(make([]byte, 0, bs.width))
	bs.bufB = bytes.NewBuffer(make([]byte, 0, bs.width))
	bs.bufA = bytes.NewBuffer(make([]byte, 0, bs.width))

	return bs
}

func syncWidth(matrix map[int][]chan int) {
	for _, column := range matrix {
		column := column
		go func() {
			var maxWidth int
			for _, ch := range column {
				if w := <-ch; w > maxWidth {
					maxWidth = w
				}
			}
			for _, ch := range column {
				ch <- maxWidth
			}
		}()
	}
}

func extractBaseFiller(f BarFiller) BarFiller {
	type wrapper interface {
		Base() BarFiller
	}
	if f, ok := f.(wrapper); ok {
		return extractBaseFiller(f.Base())
	}
	return f
}
