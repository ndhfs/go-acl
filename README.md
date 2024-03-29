# go-acl

ACL for Golang

### Example

```
var (
	ACT_RUN_FAST acl.Action = "RUN FAST"
	ACT_HIGH_JUMP acl.Action = "HIGH JUMP"

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

func main() {
	Rob.AddGroup(FootballTeam)
	Tom.AddGroup(AthletesTeam)
	Kate.AddGroup(FootballTeam, AthletesTeam)

	FootballTeam.AddAction(ACT_RUN_FAST, true)
	AthletesTeam.AddAction(ACT_HIGH_JUMP, true)
	
	for _, p := range []*Person{Rob, Tom, Kate} {
        fmt.Printf("%s can RUN FAST: %v", p.Name, p.Can(ACT_RUN_FAST))
        fmt.Printf("%s can HIGH_JUM: %v", p.Name, p.Can(ACT_HIGH_JUMP))
    }
}
```
