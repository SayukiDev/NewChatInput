package tasks

import (
	"errors"
	"github.com/orcaman/concurrent-map/v2"
)

type Tasks struct {
	runners cmap.ConcurrentMap[string, *Runner]
}

func New() *Tasks {
	return &Tasks{
		runners: cmap.New[*Runner](),
	}
}

func (t *Tasks) Add(name string, handle Handle) error {
	r := NewRunner(handle)
	_, e := t.runners.Get(name)
	if e {
		r.Close()
		return errors.New("the task is already added")
	} else {
		t.runners.Set(name, r)
	}
	r.Start()
	return nil
}

func (t *Tasks) Exist(name string) bool {
	_, ok := t.runners.Get(name)
	return ok
}

func (t *Tasks) Remove(name string) error {
	r, e := t.runners.Get(name)
	if !e {
		return errors.New("the task not exist")
	}
	t.runners.Remove(name)
	r.Close()
	return nil
}
