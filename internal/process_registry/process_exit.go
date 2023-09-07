package process_registry

import "os"

type ProcessExit struct {
	State *os.ProcessState
	Error *error
}

func NewProcessExit(state *os.ProcessState, err error) *ProcessExit {
	return &ProcessExit{State: state, Error: &err}
}
