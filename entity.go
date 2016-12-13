package acl

type Entity struct {
	abstract
	Groups []Cannable
}

func NewEntity() *Entity {
	return &Entity {
		abstract{
			can: make(map[Action]bool),
		},
		[]Cannable{},
	}
}


func (e *Entity) Can(a Action) (ok bool) {
	e.RLock()
	defer e.RUnlock()
	ok = e.abstract.Can(a)
	for _, g := range e.Groups {
		ok = ok || g.Can(a)
	}
	return
}

func (e *Entity) AddGroup(g ...Cannable) {
	e.Lock()
	defer e.Unlock()
	e.Groups = append(e.Groups, g...)
}
