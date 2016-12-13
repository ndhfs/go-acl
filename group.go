package acl

type Group struct {
	abstract
}

func NewGroup() *Group {
	return &Group {
		abstract{
			can: make(map[Action]bool),
		},
	}
}