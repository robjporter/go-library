package xordereddict

import "github.com/robjporter/go-library/xlinkedlist"

type OrderedDict struct {
	lookup map[string]*xlinkedlist.LinkedListNode
	list   *xlinkedlist.LinkedList
}

func New() *OrderedDict {
	return &OrderedDict{
		lookup: make(map[string]*xlinkedlist.LinkedListNode),
		list:   xlinkedlist.New(),
	}
}

func (d *OrderedDict) Set(key string, value interface{}) {
	if n, ok := d.lookup[key]; ok {
		d.list.Remove(n)
	}

	d.lookup[key] = d.list.Append(value)
}

func (d *OrderedDict) Get(key string) interface{} {
	if _, ok := d.lookup[key]; ok {
		return d.lookup[key].Value()
	}
	return nil
}

func (d *OrderedDict) Remove(key string) bool {
	if n, ok := d.lookup[key]; ok {
		if ok := d.list.Remove(n); !ok {
			return false
		}
		delete(d.lookup, key)
		return true
	}
	return false
}

func (d *OrderedDict) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		for v := range d.list.Iterate() {
			ch <- v.Value()
		}

		close(ch)
	}()

	return ch
}
