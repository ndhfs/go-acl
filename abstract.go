package acl

import "sync"

type Cannable interface {
	AddAction(a Action, can bool) Cannable
	Can(a Action) bool
}

type abstract struct {
	sync.RWMutex
	can map[Action]bool
}

func (e *abstract) AddAction(a Action, can bool) Cannable {
	e.Lock()
	defer e.Unlock()
	e.can[a] = can
	return e
}

func (e *abstract) Can(a Action) bool {
	e.RLock()
	defer e.RUnlock()
	b, ok := e.can[a]
	return b && ok
}
