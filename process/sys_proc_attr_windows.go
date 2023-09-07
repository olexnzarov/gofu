//go:build windows

package process

import "syscall"

// Windows-specific process attributes.
// Children spawned by the process will be created in the same group.
// When the parent exits, it should stop the children too.
func newSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	},
}
