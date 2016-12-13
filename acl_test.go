package acl_test

import (
	"git.amocrm.ru/amojo/server/core/acl"
	"testing"
)

type Person struct {
	*acl.Entity
	Name string
}

type SportTeam struct {
	*acl.Group
}

var (

	ACT_RUN_FAST acl.Action = "RUN FAST"
	ACT_HIGH_JUMP acl.Action = "HIGH_JUMP"

	Rob = &Person{
		acl.NewEntity(),
		"Rob",
	}

	Tom = &Person{
		acl.NewEntity(),
		"Tom",
	}

	Kate = &Person{
		acl.NewEntity(),
		"Kate",
	}

	FootballTeam = &SportTeam{
		acl.NewGroup(),
	}

	AthletesTeam = &SportTeam{
		acl.NewGroup(),
	}
)

func init() {
	Rob.AddGroup(FootballTeam)
	Tom.AddGroup(AthletesTeam)
	Kate.AddGroup(FootballTeam, AthletesTeam)

	FootballTeam.AddAction(ACT_RUN_FAST, true)
	AthletesTeam.AddAction(ACT_HIGH_JUMP, true)
}

func TestCan(t *testing.T) {
	for _, p := range []*Person{Rob, Tom, Kate} {
		t.Logf("%s can RUN FAST: %v", p.Name, p.Can(ACT_RUN_FAST))
		t.Logf("%s can HIGH_JUM: %v", p.Name, p.Can(ACT_HIGH_JUMP))
	}
}
