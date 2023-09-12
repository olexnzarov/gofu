//go:build linux

package process

import "syscall"

// Linux-specific process attributes.
// SIGKILL will be sent to children when the parent process exits.
func newSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}
}
