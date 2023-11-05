package processmanager

import "errors"

var (
	ErrStorageNoProcess = errors.New("storage: process not found")

	ErrInvalidProcessName          = errors.New("invalid process name")
	ErrInvalidProcessCommand       = errors.New("invalid process command")
	ErrMissingProcessConfiguration = errors.New("missing process configuration")
)
