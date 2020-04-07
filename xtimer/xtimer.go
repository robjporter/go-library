package xtimer

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/robjporter/go-library/xordereddict"
)

type Task struct {
	name string
	start time.Time
	elapsed int64
	end time.Time
}

type Timer struct {
	id              string
	tasks           *xordereddict.OrderedDict
	latestStartTime time.Time
	latestStartTask string
	totalElapsed    int64
}

func New(id string) *Timer {
	return &Timer{id: id, tasks: xordereddict.New()}
}

func (t *Timer) Start(id string) error {
	if id == "" {
		return errors.New("task name must not be empty")
	}

	if t.tasks.Get(id) != nil {
		return fmt.Errorf("Can not create new timer.  %s already exists", id)
	}

	tmp := Task{}
	tmp.name = id
	tmp.start = time.Now()

	t.tasks.Set(id, tmp)
	t.latestStartTime = time.Now()
	t.latestStartTask = id
	return nil
}

func (t *Timer) Stop(id string) error {
	if id == "" {
		return errors.New("task name must not be empty")
	}

	if t.tasks.Get(id) == nil {
		return fmt.Errorf("Can not find timer.  %s does not exist", id)
	}

	used := time.Since(t.tasks.Get(id).(Task).start).Nanoseconds()

	tmp := t.tasks.Get(id).(Task)
	tmp.elapsed = used
	tmp.end = time.Now()
	t.tasks.Set(id, tmp)

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
	if (time.Time{}) == t.tasks.Get(id).(Task).end {
		return true
	}

	return false
}

func (t *Timer) ActiveTimers() int {
	active := 0

	for value := range t.tasks.Iterate() {
		tmp := value.(Task)
		if (time.Time{}) == tmp.end {
			active++
		}
	}

	return active
}

func (t *Timer) PrettyPrint() string {
	var buf bytes.Buffer

	buf.WriteString(t.ShortSummary())

	buf.WriteString("-----------------------------------------\n")
	buf.WriteString("ms        %         Task name\n")
	buf.WriteString("-----------------------------------------\n")

	for value := range t.tasks.Iterate() {
		tmp := value.(Task)

		var elapsed int64
		if t.totalElapsed != 0 {
			elapsed = tmp.elapsed*100/t.totalElapsed
		}

		buf.WriteString(fmt.Sprintf("%-10d", tmp.elapsed/1000000))
		buf.WriteString(fmt.Sprintf("%-10s", fmt.Sprintf("%d%%", elapsed)))
		buf.WriteString(fmt.Sprintf("%s\n", tmp.name))
	}

	return buf.String()
}