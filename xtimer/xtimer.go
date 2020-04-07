package xtimer

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

type Task struct {
	name string
	start time.Time
	elapsed int64
	end time.Time
}

type Timer struct {
	id              string
	tasks           map[string]Task
	latestStartTime time.Time
	latestStartTask string
	totalElapsed    int64
}

func New(id string) *Timer {
	return &Timer{id: id, tasks: make(map[string]Task)}
}

func (t *Timer) Start(id string) error {
	if id == "" {
		return errors.New("task name must not be empty")
	}

	if t.stringInSlice(id, t.tasks) {
		return fmt.Errorf("Can not create new timer.  %s already exists", id)
	}

	tmp := Task{}
	tmp.name = id
	tmp.start = time.Now()

	t.tasks[id] = tmp
	t.latestStartTime = time.Now()
	t.latestStartTask = id
	return nil
}

func (t *Timer) Stop(id string) error {
	if id == "" {
		return errors.New("task name must not be empty")
	}

	if !t.stringInSlice(id, t.tasks) {
		return fmt.Errorf("Can not find timer.  %s does not exist", id)
	}

	used := time.Since(t.tasks[id].start).Nanoseconds()

	tmp := t.tasks[id]
	tmp.elapsed = used
	tmp.end = time.Now()
	t.tasks[id] = tmp

	t.totalElapsed += used

	return nil
}

func (t *Timer) ShortSummary() string {
	return fmt.Sprintf("Timer '"+t.id+"': running time (ms) = %d\n", t.totalElapsed/1000000)
}

func (t *Timer) LastTaskStarted() string {
	return t.latestStartTask
}

func (t *Timer) IsActive(id string) bool {
	if t.stringInSlice(id, t.tasks) {
		return true
	}

	return false
}

func (t *Timer) ActiveTimers() int {
	return len(t.tasks)
}


func (t *Timer) stringInSlice(a string, list map[string]Task) bool {
	for key, _ := range t.tasks {
		if key == a {
			return true
		}
	}
	return false
}

func (t *Timer) PrettyPrint() string {
	var buf bytes.Buffer

	buf.WriteString(t.ShortSummary())

	buf.WriteString("-----------------------------------------\n")
	buf.WriteString("ms        %         Task name\n")
	buf.WriteString("-----------------------------------------\n")

	for key, value := range t.tasks {
		var elapsed int64
		if t.totalElapsed != 0 {
			elapsed = value.elapsed*100/t.totalElapsed
		}

		buf.WriteString(fmt.Sprintf("%-10d", value.elapsed/1000000))
		buf.WriteString(fmt.Sprintf("%-10s", fmt.Sprintf("%d%%", elapsed)))
		buf.WriteString(fmt.Sprintf("%s\n", key))
	}

	return buf.String()
}

/*
func Timer(name string) string {
	name = strings.TrimSpace(name)

	if _, ok := timers[name]; ok {
		end := time.Now().Round(time.Second)
		return end.Sub(timers[name]).String()
	} else {
		timers[name] = time.Now().Round(time.Second)
	}

	return ""
}

var timers map[string]time.Time

func init() {
	timers = make(map[string]time.Time)
}

func Timer(name string) string {
	name = strings.TrimSpace(name)

	if _, ok := timers[name]; ok {
		end := time.Now().Round(time.Second)
		return end.Sub(timers[name]).String()
	} else {
		timers[name] = time.Now().Round(time.Second)
	}

	return ""
}
*/