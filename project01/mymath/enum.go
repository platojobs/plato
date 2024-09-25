package mymath

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected:
		return StateError
	case StateError:
		return StateRetrying
	case StateRetrying:
		return StateIdle
	default:
		return StateIdle
	}
}

func (ss ServerState) String() string {
	return stateName[ss]
}


func TestEnum() {

	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Print(ns2)
	fmt.Println("---------------")
    fmt.Println(ns2.String())
}

